package sender

type SendOption func(*Sender)

func SendWithAtMobiles(atMobiles []string) SendOption {
	return func(s *Sender) {
		s.at = &At{
			AtMobiles: atMobiles,
		}
	}
}

func SendWithIsAtAll(isAtAll bool) SendOption {
	return func(s *Sender) {
		s.at = &At{
			IsAtAll: isAtAll,
		}
	}
}
