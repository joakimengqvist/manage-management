/* eslint-disable @typescript-eslint/no-explicit-any */
import { createSlice } from '@reduxjs/toolkit';

export const applicationStateName = 'ReduxApplicationData';
const statename = 'application'

export const staticDataSlice = createSlice({
  name: statename,
  initialState: {
    selectedProject: {
      name: '',
      id: 0,
    },
    privileges: [],
    users: [],
    projects: []
  },
  reducers: {
    initiateApplicationData: (state : any) => {
      const savedUserState = localStorage.getItem(applicationStateName)
      if (savedUserState !== null) {
        const formattedSavedUserState = JSON.parse(savedUserState)
        return formattedSavedUserState;
      } else {
        return state;
      }
    },

    // -----------------------  |
    // -- PRIVILEGES -------    |
    // -------------------      V

    fetchPrivileges: (state : any, payload : any) => {
      state.privileges = payload.payload;
      localStorage.setItem(applicationStateName, JSON.stringify({...state, privileges: payload.payload}));
      return state;
    },
    // updatePrivilege: (state : any, payload) => {
    //  TODO
    //},
    appendPrivilege: (state : any, payload) => {
      state.privileges[state.privileges.length] = payload.payload;
      return state;
    },
    popPrivilege: (state : any, payload) => {
      state.privileges.splice(state.privileges.findIndex((p : any) => p.id === payload.payload), 1);
      return state;
    },

    // -----------------------  |
    // -- USERS ------------    |
    // -------------------      V

    fetchUsers: (state : any, payload : any) => {
      state.users = payload.payload;
      localStorage.setItem(applicationStateName, JSON.stringify({...state, users: payload.payload}));
      return state;
    },
    updateUser: (state : any, payload) => {
      const index = state.users.findIndex((u : any) => u.id === payload.payload.id);
      state.users[index] = payload.payload;
      return state;
    },
    appendUser: (state : any, payload) => {
      state.users[state.users.length] = payload.payload;
      return state;
    },
    popUser: (state : any, payload) => {
      state.users.splice(state.users.findIndex((u : any) => u.id === payload.payload), 1);
      return state;
    },

    // -----------------------  |
    // -- PROJECTS ---------    |
    // -------------------      V

    fetchProjects: (state : any, payload : any) => {
      state.projects = payload.payload;
      localStorage.setItem(applicationStateName, JSON.stringify({...state, projects: payload.payload}));
      return state;
    },
    // updateProject: (state : any, payload) => {
    //  TODO
    //},
    appendProject: (state : any, payload) => {
      state.projects[state.projects.length] = payload.payload;
      return state;
    },
    popProject: (state : any, payload) => {
      state.projects.splice(state.projects.findIndex((p : any) => p.id === payload.payload), 1);
      return state;
    },
    selectProject: (state : any, payload) => {
      state.selectedProject = payload.payload
      return state;
    },
    clearData: state => {
      const updatedData = {
          ...state,
          selectedProject: {
            name: '',
            id: 0,
          },
          privileges: [],
          users: [],
          projects: []
      }
      localStorage.setItem(applicationStateName, JSON.stringify(updatedData))
        return updatedData
    },
  }
})

export const { 
  initiateApplicationData, 
  fetchPrivileges, 
  // updatePrivilege,
  appendPrivilege,
  popPrivilege,
  fetchUsers,
  updateUser,
  appendUser,
  popUser,
  fetchProjects,
  // updateProject,
  appendProject,
  popProject,
  selectProject,
  clearData,
} = staticDataSlice.actions

export default staticDataSlice.reducer