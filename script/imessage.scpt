on run {receiver, message}
  tell application "Messages"
    send message to participant receiver of (account 1 whose service type is iMessage)
  end tell
end run