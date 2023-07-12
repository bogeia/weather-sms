on run {receiver, message}
  tell application "Messages"
    send message to buddy receiver of service "SMS"
  end tell
end run