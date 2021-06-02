import { combineReducers } from "redux";
import userLoginReducer from "./user/login/userLoginReducer";
import userRegisterReducer from "./user/register/userRegisterReducer";

const rootReducer = combineReducers({
  userRegister: userRegisterReducer,
  userLogin: userLoginReducer,
});

export default rootReducer;