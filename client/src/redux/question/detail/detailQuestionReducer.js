const initState = [];

const detailQuestionReducer = (state = initState, action) => {
  switch (action.type) {
    case "SHOW_QUESTION":
      return {
        data: action.payload.data,
      }
    case "REMOVE_DETAIL":
      return {};
    default:
      return state;
  }
}

export default detailQuestionReducer;