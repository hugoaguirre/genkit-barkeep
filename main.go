package main

import (
	"context"
	"fmt"
	"log"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/googleai"
)

const (
	FLOW_NAME    = `barCustomerFlow`
	GENAI_PROMPT = `You will tell jokes about %s, but every time the barkeep gives you a drink "
    (the barkeep will give you as a prompt: here, take this drink) you will tell only one
    joke, no more, but every time the barkeep gives you a new drink,
    you will tell another joke but a little bit more silly`
)

func main() {
	ctx := context.Background()

	// Initialize the Google AI plugin
	// GOOGLE_GENAI_API_KEY env var must be set
	if err := googleai.Init(ctx, nil); err != nil {
		log.Fatal(err)
	}

	m := googleai.Model("gemini-1.5-flash")
	if m == nil {
		log.Fatal(FLOW_NAME + ": failed to find model")
	}
	// A flow is just a function wrapper but there's a lot of fancy stuff added on top of it
	// Lets you run the function from the Genkit CLI and developer UI
	// Enables features such as deployment and observability
	// Flow responses could be simple or structured values. Genkit produces JSON schemas.
	genkit.DefineFlow(
		FLOW_NAME,
		func(ctx context.Context, input string) (string, error) {
			resp, err := ai.Generate(ctx, m,
				ai.WithConfig(&ai.GenerationCommonConfig{Temperature: 1.1}), // sober at the beginning
				ai.WithTextPrompt(fmt.Sprintf(GENAI_PROMPT, input)))
			if err != nil {
				return "", err
			}

			text := resp.Text()
			return text, nil
		})

	if err := genkit.Init(ctx, nil); err != nil {
		log.Fatal(err)
	}
}
