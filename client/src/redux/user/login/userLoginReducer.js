const initState = {
  userName: "",
  password: "",
  errorMessage: "",
  successMessage: "",
  isLoading: false,
  isLogin: false,
};

const userLoginReducer = (state = initState, action) => {
  switch (action.type) {
    case "USER_LOGIN_RESET_FORM":
      return {
        ...initState,
      };
    case "USER_LOGIN_SET_USERNAME":
      return {
        ...state,
        userName: action.payload.userName,
      };
    case "USER_LOGIN_SET_PASSWORD":
      return {
        ...state,
        password: action.payload.password,
      };
    case "USER_LOGIN_SET_ERROR_MESSAGE":
      return {
        ...state,
        errorMessage: action.payload.errorMessage,
      };
    case "USER_LOGIN_SET_SUCCESS_MESSAGE":
      return {
        ...state,
        successMessage: action.payload.successMessage,
      };
    case "USER_LOGIN_START_LOADING":
      return {
        ...state,
        isLoading: true,
      };
    case "USER_LOGIN_STOP_LOADING":
      return {
        ...state,
        isLoading: false,
      };
    case "USER_LOGIN_STATUS":
      return {
        ...state,
        isLogin: true,
      };
    case "USER_LOGOUT":
      localStorage.removeItem("accessToken");
      return {
        ...initState,
      };

    default:
      return state;
  }
};

export default userLoginReducer;