package mail

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
)

func SendEmail( ){
    fmt.Println("Sending email using AWS SES...")
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-south-1")) 
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	sesClient := ses.NewFromConfig(cfg)
	input := &ses.SendEmailInput{
		Destination: &types.Destination{
			ToAddresses: []string{"manishbiswal754@gmail.com"},
		},
		Message: &types.Message{
			Body: &types.Body{
				Text: &types.Content{
					Data: aws.String("This is a test email from Go and SES!"),
				},
			},
			Subject: &types.Content{
				Data: aws.String("Test Email"),
			},
		},
		Source: aws.String("mail@mbxd.xyz"), 
	}

	result, err := sesClient.SendEmail(context.TODO(), input)
	if err != nil {
		log.Fatalf("failed to send email: %v", err)
	}

	fmt.Println("Email sent successfully, ID:", *result.MessageId)
}
