package modules

import (
	"github.com/Greckas/SSU-Golang-252-Chat/messageService"
	"github.com/Greckas/SSU-Golang-252-Chat/server/message"
)

var changePassError = messageService.Message{
	Header:
	messageService.MessageHeader{
		Type_:   coremessage.ChangePassComm,
		Command: "Password has not changed. Please try again",
	},
	Body: nil,
}

var changePassOk = messageService.Message{
	Header:
	messageService.MessageHeader{
		Type_:   coremessage.ChangePassComm,
		Command: coremessage.StatusOk,
	},
	Body: nil,
}

var changeNicknameError = messageService.Message{
	Header:
	messageService.MessageHeader{
		Type_:   coremessage.ChangeNicknameComm,
		Command: "Nickname has not changed. Please try again",
	},
	Body: nil,
}

var changeNicknameOk = messageService.Message{
	Header:
	messageService.MessageHeader{
		Type_:   coremessage.ChangeNicknameComm,
		Command: coremessage.StatusOk,
	},
	Body: nil,
}

var changeBirthdayError = messageService.Message{
	Header:
	messageService.MessageHeader{
		Type_:   coremessage.ChangeBirthdayComm,
		Command: "Birthday has not changed. Please try again",
	},
	Body: nil,
}

var changeBirthdayOk = messageService.Message{
	Header:
	messageService.MessageHeader{
		Type_:   coremessage.ChangeBirthdayComm,
		Command: coremessage.StatusOk,
	},
	Body: nil,
}

var changeUserInfoError = messageService.Message{
	Header:
	messageService.MessageHeader{
		Type_:   coremessage.ChangeUserInfoComm,
		Command: "User info has not changed. Please try again",
	},
	Body: nil,
}

var changeUserInfoOk = messageService.Message{
	Header:
	messageService.MessageHeader{
		Type_:   coremessage.ChangeUserInfoComm,
		Command: coremessage.StatusOk,
	},
	Body: nil,
}
var blockUserError = messageService.Message{
	Header:
	messageService.MessageHeader{
		Type_:   coremessage.BlockUserComm,
		Command: "Contact has not blocked/unblocked. Please try again",
	},
	Body: nil,
}

var blockUserOk = messageService.Message{
	Header:
	messageService.MessageHeader{
		Type_:   coremessage.BlockUserComm,
		Command: coremessage.StatusOk,
	},
	Body: nil,
}
