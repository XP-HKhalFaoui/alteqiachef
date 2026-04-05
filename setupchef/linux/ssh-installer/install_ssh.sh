#!/usr/bin/env bash

set -e

if [ "$(id -u)" -ne 0 ]; then
  echo "Veuillez exécuter ce script avec sudo ou en root."
  exit 1
fi

if command -v apt >/dev/null 2>&1; then
  echo "→ Détection : Ubuntu/Debian (apt)"
  apt update
  apt install -y openssh-server   # installe le serveur SSH [web:2][web:12]
  systemctl enable ssh            # active le service au démarrage [web:1]
  systemctl start ssh             # démarre le service [web:1]
  systemctl status ssh --no-pager
elif command -v dnf >/dev/null 2>&1; then
  echo "→ Détection : Fedora (dnf)"
  dnf install -y openssh-server   # installe le serveur SSH [web:6][web:12]
  systemctl enable sshd           # active le service au démarrage [web:6][web:8]
  systemctl start sshd            # démarre le service [web:6][web:8]
  systemctl status sshd --no-pager
else
  echo "Distribution non supportée par ce script (ni apt ni dnf trouvés)."
  exit 1
fi

echo "Installation et activation du serveur SSH terminées."
