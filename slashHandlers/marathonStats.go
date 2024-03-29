package slashHandlers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"oenugs-bot/api"
	"oenugs-bot/utils"
	"strconv"
)

func MarathonStats(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// https://oengus.io/api/marathons/{marathon}/stats
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			//Flags:   1 << 6,
			Content: "Loading...",
		},
	})

	marathonId := i.ApplicationCommandData().Options[0].StringValue()

	if marathonId == "" {
		msg := "Somehow you did not supply a marathon id??"
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Content: &msg,
		})
		return
	}

	go func() {
		stats, err := api.GetMarathonStats(marathonId)

		if err != nil {
			msg := err.Error()
			s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
				Content: &msg,
			})
			return
		}

		msg := " "
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Content: &msg,
			Embeds: &[]*discordgo.MessageEmbed{
				{
					Title: fmt.Sprintf("Submission Stats for **%s**", marathonId),
					Fields: []*discordgo.MessageEmbedField{
						{
							Name:  "Total Submissions",
							Value: strconv.Itoa(stats.SubmissionCount),
						},
						{
							Name:  "Total Runners",
							Value: strconv.Itoa(stats.RunnerCount),
						},
						{
							Name:  "Total Length",
							Value: utils.ParseAndMakeDurationPretty(stats.TotalLength),
						},
						{
							Name:  "Average Estimate",
							Value: utils.ParseAndMakeDurationPretty(stats.AverageEstimate),
						},
					},
				},
			},
		})
	}()
}
