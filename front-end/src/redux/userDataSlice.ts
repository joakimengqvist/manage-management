/* eslint-disable @typescript-eslint/no-explicit-any */
import { createSlice } from '@reduxjs/toolkit';

export const userSlice = createSlice({
  name: 'user',
  initialState: {
    authenticated: false,
    email: '',
    firstName: '',
    lastName: '',
    id: "",
    privileges: [],
    projects: [],
  },
  reducers: {
    initiateUser: (state : any) => {
      const savedUserState = localStorage.getItem("userReduxData")
      if (savedUserState !== null) {
        const formattedSavedUserState = JSON.parse(savedUserState)
        return formattedSavedUserState;
      } else {
        return state;
      }
    },
    authenticate: (state : any, payload : any) => {
        const data = payload.payload.data;
        const updatedData = {
          ...state,
          authenticated: true,
          email: data.email,
          firstName: data.first_name,
          lastName: data.last_name,
          id: data.id,
          privileges: data.privileges,
          projects: data.projects,
      }
        localStorage.setItem("userReduxData", JSON.stringify(updatedData))
        return updatedData
    },
    logout: state => {
      const updatedData = {
          ...state,
          authenticated: false,
          email: "",
          firstName: "",
          lastName: "",
          id: "",
          projects: [],
      }
      localStorage.setItem("userReduxData", JSON.stringify(updatedData))
        return updatedData
    },
  }
})

// Action creators are generated for each case reducer function
export const { initiateUser, authenticate, logout } = userSlice.actions

export default userSlice.reducer