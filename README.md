# [WIP] A barkeep that gets GoogleAI drunk

This demo was created by following [Genkit Golang documentation](https://firebase.google.com/docs/genkit-go/get-started-go)

*Note: Ideally, this app should allow users (barkeep) to keep a conversation with `Google Gemini AI` and sip drinks to it and the reward will be super silly jokes every time a drink gets served to the AI*

# Dependencies

- Node.js 20 or later
- Go 1.22 or later

# How to run it

```shell
  ## Install genkit
  $ npm i -g genkit

  ## On the project folder, run the following commands
  $ genkit init --model googleai

  ## IMPORTANT: provide your Google API key
  export GOOGLE_GENAI_API_KEY=foobar

  ## Run the app
  $ genkit start
```
