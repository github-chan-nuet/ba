package services

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"phishing_backend/internal/domain_model"
	"testing"
)

type FeatureValue = domain_model.PhishingSimulationRecognitionFeatureValue
type RecognitionFeature = domain_model.PhishingSimulationRecognitionFeature

func TestParseTemplate(t *testing.T) {
	// given
	sut := PhishingEmailGenerationServiceImpl{}
	run := domain_model.PhishingSimulationRun{}
	placeholderHandlers["A"] = createPlaceholderHandler("a")
	placeholderHandlers["B"] = createPlaceholderHandler("b")
	placeholderHandlers["C"] = createPlaceholderHandler("c")
	placeholderHandlers["D"] = func(arg string, _ *domain_model.PhishingSimulationRun) (string, error) {
		return "{{E}}", nil
	}
	placeholderHandlers["E"] = func(arg string, _ *domain_model.PhishingSimulationRun) (string, error) {
		return "e", nil
	}

	tests := []struct {
		name string
		tpl  string
		want string
	}{
		{
			name: "No template",
			tpl:  "test test",
			want: "test test",
		},
		{
			name: "Simple templating",
			tpl:  "{{A}}",
			want: "a",
		},
		{
			name: "Templating with argument",
			tpl:  "{{A hallo}}",
			want: "a arg:hallo",
		},
		{
			name: "Templating with argument and many spaces",
			tpl:  "{{A     hallo}}",
			want: "a arg:    hallo",
		},
		{
			name: "Templating with argument with same placeholder name",
			tpl:  "{{A A}}",
			want: "a arg:A",
		},
		{
			name: "Nested templating with arg",
			tpl:  "{{A {{B}}}}",
			want: "a arg:b",
		},
		{
			name: "Double nested templating with arg",
			tpl:  "{{A {{B {{C}}}}}}",
			want: "a arg:b arg:c",
		},
		{
			name: "Several templates",
			tpl:  "{{A}} {{B}} {{C}}",
			want: "a b c",
		},
		{
			name: "Several templates with arguments",
			tpl:  "{{A 123}} {{B 456}} {{C 789}}",
			want: "a arg:123 b arg:456 c arg:789",
		},
		{
			name: "Several templates some with arguments",
			tpl:  "{{A}} {{B 456}} {{C}}",
			want: "a b arg:456 c",
		},
		{
			name: "Template with surrounding text",
			tpl:  "Hallo{{A}}Welt",
			want: "HalloaWelt",
		},
		{
			name: "Recursive template",
			tpl:  "{{D}}",
			want: "e",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// when
			result, err := sut.parseTemplate(tt.tpl, &run)

			// then
			assert.NoError(t, err)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestGenerateEmail(t *testing.T) {
	// given
	sut := PhishingEmailGenerationServiceImpl{}
	placeholderHandlers["A"] = createPlaceholderHandler("a")
	placeholderHandlers["B"] = createPlaceholderHandler("b")
	run := domain_model.PhishingSimulationRun{
		ID:                       uuid.New(),
		User:                     &domain_model.User{Email: "a@gmail.com"},
		RecognitionFeatureValues: []FeatureValue{{RecognitionFeature: &RecognitionFeature{Name: "Domain"}, Value: "AH"}},
		Template: &domain_model.PhishingSimulationContentTemplate{
			Subject: "Hallo {{A}}",
			Content: "Goodbye {{B}}",
		},
	}

	// when
	mail := sut.GenerateEmail(&run)

	// then
	assert.Equal(t, "Hallo a", mail.Subject)
	assert.Equal(t, "Goodbye b", mail.Content)
	assert.Equal(t, "info@AH", mail.Sender)
	assert.Equal(t, run.User.Email, mail.Recipient)
}

func createPlaceholderHandler(placeHolder string) PlaceholderHandler {
	return func(arg string, _ *domain_model.PhishingSimulationRun) (string, error) {
		if arg == "" {
			return placeHolder, nil
		}
		return placeHolder + " arg:" + arg, nil
	}
}
