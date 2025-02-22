package entities

type Word struct {
    ID            int    `json:"id"`
    Word          string `json:"word"`
    Lang         string `json:"lang"`
    PartOfSpeech string `json:"part_of_speech"`
    Transcription string `json:"transcription"`
}

type WordTranslation struct {
    MeaningID     int `json:"meaning_id"`
    TranslationID int `json:"translation_id"`
}

type WordGroup struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    LangFrom string `json:"lang_from"`
    LangTo   string `json:"lang_to"`
    UserID   int    `json:"user_id"`
    Status   string `json:"status"`
}

type WordGroupWord struct {
    GroupID       int    `json:"group_id"`
    WordMeaningID int    `json:"word_meaning_id"`
    Level         string `json:"level"`
    Status        string `json:"status"`
    LastRight     string `json:"last_right"`
}