package get

import (
	"github.com/google/uuid"
	"hedgehog-forms/model/form/generated"
)

//        Map<UUID, @NotNull UUID> singleChoiceAnswers,
//
//        @NotNull(message = "Не может быть null, должно быть хотя бы пустым")
//        @Schema(description = "Мапа для ответов на вопросы со множественным выбором. " +
//                "Должна быть хотя бы пустой, но не null. Ключ - ID вопроса, значение - список ID ответов",
//                requiredMode = REQUIRED)
//        Map<UUID, @NotEmpty List<@NotNull UUID>> multipleChoiceAnswers,
//
//        @NotNull(message = "Не может быть null, должно быть хотя бы пустым")
//        @Schema(description = "Мапа для ответов на вопросы с текстовым вводом. " +
//                "Должна быть хотя бы пустой, но не null. Ключ - ID вопроса, значение - текст ответа",
//                requiredMode = REQUIRED)
//        Map<UUID, @NotEmpty String> textAnswers,
//
//        @NotNull(message = "Не может быть null, должно быть хотя бы пустым")
//        @Schema(description = "Мапа для ответов на вопросы с сопоставлением. Ключ - ID вопроса, значение - список, " +
//                "где элементы списка - пары. Первый элемент пары - ID термина, второй - ID определения",
//                requiredMode = REQUIRED)
//        Map<UUID, List<Pair<UUID, UUID>>> matchingAnswers

type AnswerDto struct {
	SingleChoice   map[uuid.UUID]uuid.UUID
	MultipleChoice map[uuid.UUID][]uuid.UUID
	TextInput      map[uuid.UUID]string
	Matching       map[uuid.UUID][]generated.EnteredMatchingPair
}
