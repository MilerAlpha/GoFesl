package theater

import (
	"github.com/SpencerSharkey/GoFesl/GameSpy"
	"github.com/SpencerSharkey/GoFesl/log"
)

// EGRS - SERVER sent up, tell us if client is 'allowed' to join
func (tM *TheaterManager) EGRS(event GameSpy.EventClientFESLCommand) {
	if !event.Client.IsActive {
		return
	}

	_, err := tM.stmtGameIncreaseJoining.Exec(event.Command.Message["GID"], Shard)
	if err != nil {
		log.Panicln(err)
	}

	answer := make(map[string]string)
	answer["TID"] = event.Command.Message["TID"]
	event.Client.WriteFESL("EGRS", answer, 0x0)
}
