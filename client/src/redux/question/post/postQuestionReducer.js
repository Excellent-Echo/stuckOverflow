const initState = {
  title: "",
  content: "",
  categoryID: "",
  errorMessage: "",
  successMessage: "",
  isLoading: false,
};

const postQuestionReducer = (state = initState, action) => {
  switch (action.type) {
    case "ASK_QUESTION_RESET_FORM":
      return {
        ...initState,
      };
    case "ASK_QUESTION_SET_TITLE":
      return {
        ...state,
        title: action.payload.title,
      };
    case "ASK_QUESTION_SET_CONTENT":
      return {
        ...state,
        content: action.payload.content,
      };
    case "ASK_QUESTION_SET_CATEGORY":
      return {
        ...state,
        categoryID: action.payload.categoryID,
      };
    case "ASK_QUESTION_SET_ERROR_MESSAGE":
      return {
        ...state,
        errorMessage: action.payload.errorMessage,
      };
    // case "ASK_QUESTION_SET_SUCCESS_MESSAGE":
    //   return {
    //     ...state,
    //     successMessage: action.payload.successMessage,
    //   };
    case "ASK_QUESTION_START_LOADING":
      return {
        ...state,
        isLoading: true,
      };
    case "ASK_QUESTION_STOP_LOADING":
      return {
        ...state,
        isLoading: false,
      };

    default:
      return state;
  }
};

export default postQuestionReducer;