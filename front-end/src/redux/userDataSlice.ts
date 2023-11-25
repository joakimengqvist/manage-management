/* eslint-disable @typescript-eslint/no-explicit-any */
import { createSlice } from '@reduxjs/toolkit';

const initialState = {
  authenticated: false,
  email: '',
  firstName: '',
  lastName: '',
  id: "",
  privileges: [],
  projects: [],
  settings: {
    dark_theme: false,
    compact_ui: false,
  }
};

export const userSlice = createSlice({
  name: 'user',
  initialState: initialState,
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
    authenticate: (state : any, payload) => {
        const data = payload.payload;
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
    fetchUserSettings: (state : any, payload) => {
      const settings = {
        compact_ui: payload.payload.compact_ui,
        dark_theme: payload.payload.dark_theme,
       } || initialState.settings;
      state.settings = settings;
      localStorage.setItem("userReduxData", JSON.stringify({ ...state, settings: settings }));
      return state;
    },
    updateDarkTheme: (state : any, payload) => {
      const darkTheme = payload.payload;
      state.settings.dark_theme = darkTheme;
      localStorage.setItem("userReduxData", JSON.stringify({ ...state, settings: { ...state.settings, dark_theme: darkTheme } }));
      return state;
    },
    updateCompactUI: (state : any, payload) => {
      const compactUI = payload.payload;
      state.settings.compact_ui = compactUI;
      localStorage.setItem("userReduxData", JSON.stringify({ ...state, settings: { ...state.settings, compact_ui: compactUI } }));
      return state;
    },
    logout: state => {
      const updatedData = {
          ...state,
          authenticated: false,
          email: "",
          firstName: "",
          lastName: "",
          id: "",
          privileges: [],
          projects: [],
          settings: {
            dark_theme: state.settings.dark_theme,
            compact_ui: state.settings.compact_ui,
          }

      }
      localStorage.setItem("userReduxData", JSON.stringify(updatedData))
        return updatedData
    },
  }
})

// Action creators are generated for each case reducer function
export const { initiateUser, authenticate, fetchUserSettings, updateDarkTheme, updateCompactUI, logout } = userSlice.actions

export default userSlice.reducer