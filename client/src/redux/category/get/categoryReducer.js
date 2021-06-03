const initState = [];

const categoryReducer = (state = initState, action) => {
  switch (action.type) {
    case "SEARCH_CATEGORY":
      return {
        data: action.payload.list,
      };
    default:
      return state;
  }
}

export default categoryReducer;