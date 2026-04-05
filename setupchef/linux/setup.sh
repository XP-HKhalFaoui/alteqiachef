#!/bin/bash

# AlteqiaChef Setup Script for Linux (Fedora/RHEL/CentOS)
# This script automatically installs Docker and starts the POS system
# Usage: ./setup.sh [action]
# Actions: setup, start, stop, restart, status, logs, clean, help

set -e

# Color definitions
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Detect Linux distribution
detect_distro() {
    if [ -f /etc/os-release ]; then
        . /etc/os-release
        echo "$ID"
    else
        echo "unknown"
    fi
}

DISTRO=$(detect_distro)

# Function to print colored output
print_header() {
    echo -e "${BLUE}╔════════════════════════════════════════════════════════╗${NC}"
    echo -e "${BLUE}║ $1${NC}"
    echo -e "${BLUE}╚════════════════════════════════════════════════════════╝${NC}"
    echo ""
}

print_info() {
    echo -e "${BLUE}[*]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[✓]${NC} $1"
}

print_error() {
    echo -e "${RED}[✗]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[!]${NC} $1"
}

# Check if running as root (for package installation)
check_sudo() {
    if [ "$EUID" -ne 0 ]; then
        print_warning "This operation requires administrator privileges"
        print_info "Please provide your password when prompted"
        sudo -v
        if [ $? -ne 0 ]; then
            print_error "Failed to obtain sudo privileges"
            exit 1
        fi
    fi
}

# Install Docker
install_docker() {
    print_header "Docker Installation"
    
    print_info "Installing Docker for $DISTRO..."
    check_sudo
    
    case "$DISTRO" in
        fedora|rhel|centos)
            print_info "Installing Docker via DNF..."
            sudo dnf install -y dnf-plugins-core
            sudo dnf config-manager --add-repo https://download.docker.com/linux/fedora/docker-ce.repo
            sudo dnf install -y docker-ce docker-ce-cli containerd.io docker-compose-plugin
            ;;
        debian|ubuntu|zorin)
            print_info "Installing Docker via APT..."
            sudo apt-get update
            sudo apt-get install -y docker.io docker-compose
            ;;
        *)
            print_error "Unsupported distribution: $DISTRO"
            print_info "Supported: fedora, rhel, centos, debian, ubuntu, zorin"
            exit 1
            ;;
    esac
    
    if [ $? -eq 0 ]; then
        print_success "Docker installed successfully"
    else
        print_error "Failed to install Docker"
        exit 1
    fi
}

# Docker Compose wrapper function - supports both old and new syntax
# Uses the docker-compose.yml from parent directory
dc() {
    local compose_cmd
    local compose_file="docker-compose.yml"
    
    if command -v docker-compose &> /dev/null; then
        compose_cmd="docker-compose"
    else
        compose_cmd="docker compose"
    fi
    
    # Try running normally first
    if ! $compose_cmd -f "$compose_file" "$@" 2>/dev/null; then
        # If failed, try with sudo
        print_warning "Attempting docker command with elevated privileges..."
        sudo $compose_cmd -f "$compose_file" "$@"
    fi
}

# Check if Docker is installed
check_docker() {
    print_info "Checking Docker installation..."
    
    if command -v docker &> /dev/null; then
        local docker_version=$(docker --version)
        print_success "Docker found: $docker_version"
        return 0
    else
        print_error "Docker is not installed"
        return 1
    fi
}

# Verify Docker Compose
verify_docker_compose() {
    print_info "Checking Docker Compose..."
    
    if command -v docker-compose &> /dev/null; then
        local compose_version=$(docker-compose --version)
        print_success "Docker Compose found: $compose_version"
        return 0
    elif docker compose version &> /dev/null; then
        print_success "Docker Compose plugin found"
        return 0
    else
        print_error "Docker Compose is not installed"
        return 1
    fi
}

# Check if Docker daemon is running
check_docker_daemon() {
    print_info "Checking Docker daemon..."
    
    if systemctl is-active --quiet docker; then
        print_success "Docker daemon is running"
        return 0
    else
        print_warning "Docker daemon is not running"
        print_info "Starting Docker daemon..."
        sudo systemctl start docker
        
        if [ $? -eq 0 ]; then
            print_success "Docker daemon started"
            # Enable auto-start
            sudo systemctl enable docker
            return 0
        else
            print_error "Failed to start Docker daemon"
            return 1
        fi
    fi
}

# Add current user to docker group
add_user_to_docker() {
    print_info "Configuring Docker permissions..."
    
    if groups "$USER" | grep -q docker; then
        print_success "User is in docker group"
        # Test if docker actually works without sudo
        if docker ps &>/dev/null; then
            print_success "Docker commands work without sudo"
            return 0
        else
            print_warning "User in docker group but socket permission issue detected"
            print_info "Docker commands will use sudo when needed"
            return 0
        fi
    else
        print_warning "Current user not in docker group"
        print_info "Adding user to docker group..."
        sudo usermod -aG docker "$USER"
        print_success "User added to docker group"
        print_info "Docker commands will use sudo when needed (group change requires new session)"
        return 0
    fi
}

# Check system ports
check_ports() {
    print_info "Checking required ports..."
    
    local ports=(3000 8080 5432)
    local ports_in_use=()
    
    for port in "${ports[@]}"; do
        if netstat -tuln 2>/dev/null | grep -q ":$port "; then
            ports_in_use+=($port)
        fi
    done
    
    if [ ${#ports_in_use[@]} -gt 0 ]; then
        print_warning "Port(s) already in use: ${ports_in_use[@]}"
        print_info "These may be from previous containers or other applications"
        print_info "Run './setup.sh stop' to stop existing containers"
    else
        print_success "All required ports are available"
    fi
}

# Pull Docker images
pull_images() {
    print_header "Pulling Docker Images"
    
    print_info "Pulling latest Docker images..."
    if dc pull; then
        print_success "Docker images pulled successfully"
        return 0
    else
        print_error "Failed to pull Docker images"
        return 1
    fi
}

# Start containers
start_containers() {
    print_header "Starting Containers"
    
    print_info "Starting Docker containers..."
    if dc up -d; then
        print_success "Containers started successfully"
        return 0
    else
        print_error "Failed to start containers"
        return 1
    fi
}

# Show service status
show_status() {
    print_header "Service Status"
    
    dc ps
    echo ""
    
    print_info "Service Health:"
    echo ""
    
    # Check Frontend
    if curl -s -o /dev/null -w "%{http_code}" http://localhost:3000 | grep -q "200\|301\|302"; then
        print_success "Frontend (3000): Running"
    else
        print_error "Frontend (3000): Not responding"
    fi
    
    # Check Backend
    if curl -s -o /dev/null -w "%{http_code}" http://localhost:8080/api/v1/health | grep -q "200"; then
        print_success "Backend (8080): Running"
    else
        print_error "Backend (8080): Not responding"
    fi
    
    # Check Database
    if dc exec -T postgres pg_isready -U postgres &> /dev/null; then
        print_success "Database (5432): Running"
    else
        print_error "Database (5432): Not responding"
    fi
    
    echo ""
}

# Show logs
show_logs() {
    print_header "Application Logs"
    
    print_warning "Press Ctrl+C to stop viewing logs"
    echo ""
    
    dc logs -f
}

# Stop containers
stop_containers() {
    print_header "Stopping Containers"
    
    print_info "Stopping Docker containers..."
    dc down
    
    print_success "Containers stopped"
}

# Restart containers
restart_containers() {
    print_header "Restarting Containers"
    
    print_info "Restarting Docker containers..."
    dc restart
    
    echo ""
    dc ps
    echo ""
}

# Clean everything
clean_all() {
    print_header "Cleanup"
    
    print_warning "This will:"
    print_warning "  • Stop all containers"
    print_warning "  • Remove containers and networks"
    print_warning "  • Remove volumes (DATABASE WILL BE LOST!)"
    echo ""
    
    read -p "Are you sure? Type 'yes' to confirm: " confirm
    
    if [ "$confirm" = "yes" ]; then
        print_info "Removing containers and volumes..."
        dc down -v
        print_success "Cleanup complete"
    else
        print_error "Cleanup cancelled"
    fi
}

# Complete setup
setup_application() {
    print_header "AlteqiaChef Setup - Fedora/Linux"
    
    # Check and install Docker if needed
    if ! check_docker; then
        read -p "Docker not found. Install Docker? (y/n): " install_choice
        if [ "$install_choice" = "y" ] || [ "$install_choice" = "yes" ]; then
            install_docker
        else
            print_error "Docker is required. Exiting."
            exit 1
        fi
    fi
    
    # Verify Docker Compose
    if ! verify_docker_compose; then
        print_error "Docker Compose is required"
        exit 1
    fi
    
    # Start Docker daemon
    if ! check_docker_daemon; then
        print_error "Failed to start Docker daemon"
        exit 1
    fi
    
    # Add user to docker group
    add_user_to_docker
    
    # Check ports
    check_ports
    
    echo ""
    
    # Pull images
    if ! pull_images; then
        print_error "Failed to pull images"
        exit 1
    fi
    
    echo ""
    
    # Start containers
    if ! start_containers; then
        print_error "Failed to start containers"
        exit 1
    fi
    
    # Show status
    sleep 3
    show_status
    
    # Show success message
    print_header "Setup Complete! 🎉"
    
    print_success "AlteqiaChef is now running!"
    echo ""
    print_info "Access your application:"
    print_info "  Frontend: http://localhost:3000"
    print_info "  Backend:  http://localhost:8080"
    print_info "  Database: localhost:5432"
    echo ""
    print_info "Next steps:"
    print_info "  1. Open http://localhost:3000 in your browser"
    print_info "  2. Create an admin user ( ./create-admin.sh)"
    print_info "  3. Start managing your restaurant!"
    echo ""
}

# Show help
show_help() {
    print_header "AlteqiaChef Setup Script Help"
    
    echo "USAGE:"
    echo "  ./setup.sh [action]"
    echo ""
    echo "ACTIONS:"
    echo "  setup      : Complete setup (install Docker, pull images, start containers)"
    echo "  start      : Start containers"
    echo "  stop       : Stop containers"
    echo "  restart    : Restart all containers"
    echo "  status     : Show service status"
    echo "  logs       : Stream logs from all services"
    echo "  clean      : Remove containers and volumes (WARNING: loses data)"
    echo "  help       : Show this help message"
    echo ""
    echo "EXAMPLES:"
    echo "  ./setup.sh setup"
    echo "  ./setup.sh status"
    echo "  ./setup.sh logs"
    echo ""
    echo "LOCATIONS:"
    echo "  Frontend: http://localhost:3000"
    echo "  Backend:  http://localhost:8080"
    echo "  Database: localhost:5432"
    echo ""
}

# Main execution
ACTION="${1:-help}"

case "$ACTION" in
    setup)
        setup_application
        ;;
    start)
         dc up -d
        show_status
        ;;
    stop)
        stop_containers
        ;;
    restart)
        restart_containers
        ;;
    status)
        show_status
        ;;
    logs)
        show_logs
        ;;
    clean)
        clean_all
        ;;
    help)
        show_help
        ;;
    *)
        print_error "Unknown action: $ACTION"
        show_help
        exit 1
        ;;
esac
