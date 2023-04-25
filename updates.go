package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handle_update(update tgbotapi.Update, bot *tgbotapi.BotAPI) error {
	if update.Message == nil { // ignore any non-Message updates
		return nil
	}
	if !update.Message.IsCommand() { // ignore any non-command Messages
		if len(update.Message.Photo) > 0 {
			updateUserImage(update.Message.From.UserName)
		}
		return nil
	}
	if err := handle_command_update(update, bot); err != nil {
		return err
	}
	return nil
}

func handle_command_update(update tgbotapi.Update, bot *tgbotapi.BotAPI) error {
	var msg tgbotapi.MessageConfig
	switch update.Message.Command() {
	case "status":
		msg = handle_status_command(update)
	case "show_users":
		msg = handle_show_users_command(update)
	case "workout":
		msg = handle_workout_command(update)
	case "admin_delete_last":
		msg = handle_admin_delete_last_command(update, bot)
	default:
		msg.Text = "Unknown command"
	}
	if _, err := bot.Send(msg); err != nil {
		return err
	}
	return nil
}