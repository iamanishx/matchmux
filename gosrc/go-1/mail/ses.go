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

func SendEmail(mail string, otp string) {
    fmt.Println("Sending email using AWS SES...")
    cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-south-1"))
    if err != nil {
        log.Fatalf("unable to load SDK config, %v", err)
    }

    sesClient := ses.NewFromConfig(cfg)
    input := &ses.SendEmailInput{
        Destination: &types.Destination{
            ToAddresses: []string{mail}, 
        },
        Message: &types.Message{
            Body: &types.Body{
                Text: &types.Content{
                    Data: aws.String(fmt.Sprintf("Your verification code is: %s\n\nThis code will expire in 10 minutes.", otp)),
                },
            },
            Subject: &types.Content{
                Data: aws.String("Email Verification Code"),
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
