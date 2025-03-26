#!/bin/bash

# 환경 설정
SERVER_USER=${1:-"root"}       # 첫 번째 인자로 사용자 계정 지정 (없으면 기본값 root)
SERVER_IP=${2:-"192.168.68.7"} # 두 번째 인자로 서버 IP 지정 (없으면 기본값)
SERVER_PORT=${3:-"22"}         # 세 번째 인자로 SSH 포트 지정 (없으면 기본값 22)
SSH_KEY_PATH=${4:-"/Users/sanguklee/Desktop/git/2025/Personal/system-Info-collector/script/ssh-key/id_rsa"}
TARGET_DIR="/usr/local/bin"
SERVICE_NAME="system-agent"
SERVICE_FILE="/etc/systemd/system/${SERVICE_NAME}.service"
BINARY_FILE="$TARGET_DIR/system-agent"

# 1. 서비스 중지
echo "🛑 서비스 중지 중..."
ssh -i "$SSH_KEY_PATH" -p $SERVER_PORT $SERVER_USER@$SERVER_IP "sudo systemctl stop $SERVICE_NAME"

# 2. 서비스 비활성화
echo "🚫 서비스 비활성화..."
ssh -i "$SSH_KEY_PATH" -p $SERVER_PORT $SERVER_USER@$SERVER_IP "sudo systemctl disable $SERVICE_NAME"

# 3. 서비스 파일 삭제
echo "🗑 서비스 파일 삭제..."
ssh -i "$SSH_KEY_PATH" -p $SERVER_PORT $SERVER_USER@$SERVER_IP "sudo rm -f $SERVICE_FILE"

# 4. 실행 파일 삭제
echo "🗑 실행 파일 삭제..."
ssh -i "$SSH_KEY_PATH" -p $SERVER_PORT $SERVER_USER@$SERVER_IP "sudo rm -f $BINARY_FILE"

# 5. systemd 데몬 리로드
echo "🔄 systemd 데몬 리로드..."
ssh -i "$SSH_KEY_PATH" -p $SERVER_PORT $SERVER_USER@$SERVER_IP "sudo systemctl daemon-reload"

# 6. 로그 정리 (잔여 로그 제거)
echo "🧹 서비스 로그 삭제..."
ssh -i "$SSH_KEY_PATH" -p $SERVER_PORT $SERVER_USER@$SERVER_IP "sudo journalctl --vacuum-time=1s --unit=$SERVICE_NAME"

echo "✅ 서비스 삭제 완료!"
