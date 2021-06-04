const initState = [];

const searchQuestionReducer = (state = initState, action) => {
  switch (action.type) {
    case "SEARCH_QUESTIONS":
      return {
        loaded: true,
        data: action.payload.list,
      }
    default:
      return state;
  }
}

export default searchQuestionReducer;