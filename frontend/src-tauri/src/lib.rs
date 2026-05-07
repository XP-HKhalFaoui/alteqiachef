use std::sync::Mutex;
use tauri::{
    menu::{Menu, MenuItem},
    tray::TrayIconBuilder,
    AppHandle, Manager, RunEvent,
};
use tauri_plugin_shell::process::CommandChild;
use tauri_plugin_shell::ShellExt;

struct SidecarHandle(Mutex<Option<CommandChild>>);

pub fn run() {
    tauri::Builder::default()
        .plugin(tauri_plugin_shell::init())
        .manage(SidecarHandle(Mutex::new(None)))
        .setup(|app| {
            // Resolve DB path in app data directory
            let app_data = app
                .path()
                .app_data_dir()
                .expect("failed to resolve app data dir");
            std::fs::create_dir_all(&app_data).ok();
            let db_path = app_data.join("pos.db");

            // Spawn Go sidecar
            let sidecar_cmd = app
                .shell()
                .sidecar("pos-backend")
                .expect("pos-backend sidecar not found")
                .env("ALTEQIA_DB_PATH", db_path.to_string_lossy().to_string())
                .env("PORT", "17432");

            let (_, child) = sidecar_cmd.spawn().expect("failed to spawn pos-backend");
            *app.state::<SidecarHandle>().0.lock().unwrap() = Some(child);

            // System tray
            let show = MenuItem::with_id(app, "show", "Show Window", true, None::<&str>)?;
            let quit = MenuItem::with_id(app, "quit", "Quit AlteqiaChef", true, None::<&str>)?;
            let menu = Menu::with_items(app, &[&show, &quit])?;

            TrayIconBuilder::new()
                .menu(&menu)
                .on_menu_event(|app, event| match event.id.as_ref() {
                    "show" => {
                        if let Some(window) = app.get_webview_window("main") {
                            window.show().ok();
                            window.set_focus().ok();
                        }
                    }
                    "quit" => kill_sidecar_and_exit(app),
                    _ => {}
                })
                .build(app)?;

            Ok(())
        })
        .on_window_event(|window, event| {
            if let tauri::WindowEvent::CloseRequested { api, .. } = event {
                api.prevent_close();
                window.hide().ok();
            }
        })
        .build(tauri::generate_context!())
        .expect("error building tauri app")
        .run(|app, event| {
            if let RunEvent::ExitRequested { .. } = event {
                kill_sidecar_and_exit(app);
            }
        });
}

fn kill_sidecar_and_exit(app: &AppHandle) {
    if let Ok(mut guard) = app.state::<SidecarHandle>().0.lock() {
        if let Some(child) = guard.take() {
            child.kill().ok();
        }
    }
    app.exit(0);
}
