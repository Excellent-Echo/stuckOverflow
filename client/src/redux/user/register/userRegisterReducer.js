const initState = {
  userName: "",
  email: "",
  password: "",
  errorMessage: "",
  successMessage: "",
  isLoading: false,
};

const userRegisterReducer = (state = initState, action) => {
  switch (action.type) {
    case "USER_REGISTER_RESET_FORM":
      return {
        ...initState,
      };
    case "USER_REGISTER_SET_USERNAME":
      return {
        ...state,
        userName: action.payload.userName,
      };
    case "USER_REGISTER_SET_EMAIL":
      return {
        ...state,
        email: action.payload.email,
      };
    case "USER_REGISTER_SET_PASSWORD":
      return {
        ...state,
        password: action.payload.password,
      };
    case "USER_REGISTER_SET_ERROR_MESSAGE":
      return {
        ...state,
        errorMessage: action.payload.errorMessage,
      };
    case "USER_REGISTER_SET_SUCCESS_MESSAGE":
      return {
        ...state,
        successMessage: action.payload.successMessage,
      };
    case "USER_REGISTER_START_LOADING":
      return {
        ...state,
        isLoading: true,
      };
    case "USER_REGISTER_STOP_LOADING":
      return {
        ...state,
        isLoading: false,
      };

    default:
      return state;
  }
};

export default userRegisterReducer;