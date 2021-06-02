import { combineReducers } from "redux";
import postQuestionReducer from "./question/post/postQuestionReducer";
import questionReducer from "./question/questionReducer";
import userLoginReducer from "./user/login/userLoginReducer";
import userRegisterReducer from "./user/register/userRegisterReducer";

const rootReducer = combineReducers({
  userRegister: userRegisterReducer,
  userLogin: userLoginReducer,
  question: questionReducer,
  postQuestion: postQuestionReducer,
});

export default rootReducer;