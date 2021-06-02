const initState = [];

const questionReducer = (state = initState, action) => {
  switch (action.type) {
    case "SHOW_ALL_QUESTIONS":
      return {
        loaded: true,
        data: action.payload.list,
      }
    default:
      return state;
  }
}

export default questionReducer;