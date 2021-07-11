package grug

func init() {
	AllActions = append(AllActions, []Action{
		{
			// Takes the channel ID of the command message and returns it
			Name: "GetCommandChannelID",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				return g.CurrentCommand.ChannelID, nil
			},
		},
		{
			// Takes the guild ID of the command message and returns it
			Name: "GetCommandGuildID",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				return g.CurrentCommand.GuildID, nil
			},
		},
		{
			// Takes the user of the command message and returns it
			Name: "GetCommandUser",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				return g.CurrentCommand.Member.User, nil
			},
		},
	}...)
}
