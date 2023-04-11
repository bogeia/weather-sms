on run {receiver, message}
  tell application "Messages"
    send message to participant receiver of service "SMS"
  end tell
end run