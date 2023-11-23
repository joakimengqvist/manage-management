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
    projects: [],
    subProjects: [],
    products: [],
    invoices: [],
    invoiceItems: [],
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
      const privileges = payload.payload || [];
      state.privileges = privileges;
      localStorage.setItem(applicationStateName, JSON.stringify({ ...state, privileges: privileges }));
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
      const users = payload.payload || [];
      state.users = users;
      localStorage.setItem(applicationStateName, JSON.stringify({ ...state, users: users }));
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
      const projects = payload.payload || [];
      state.projects = projects;
      localStorage.setItem(applicationStateName, JSON.stringify({ ...state, projects: projects }));
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

    // -------------------------  |
    // -- SUB PROJECTS -------    |
    // -------------------        V

    fetchSubProjects: (state : any, payload : any) => {
      const payloadSubProjects = payload.payload || [];
      state.subProjects = payloadSubProjects;
      localStorage.setItem(applicationStateName, JSON.stringify({ ...state, subProjects: payloadSubProjects }));
      return state;
    },
    // updatesubProject: (state : any, payload) => {
    //  TODO
    //},
    appendSubProject: (state : any, payload) => {
      state.subProjects[state.subProjects.length] = payload.payload;
      return state;
    },
    popSubProject: (state : any, payload) => {
      state.subProjects.splice(state.subProjects.findIndex((p : any) => p.id === payload.payload), 1);
      return state;
    },

    // -----------------------  |
    // -- EXTERNAL COMPANIES    |
    // -------------------      V

    fetchExternalCompanies: (state : any, payload) => {
      const payloadExternalCompanies = payload.payload || [];
      state.externalCompanies = payloadExternalCompanies;
      localStorage.setItem(applicationStateName, JSON.stringify({ ...state, externalCompanies: payloadExternalCompanies }));
      return state;
    },

    // -----------------------  |
    // -- EXTERNAL COMPANIES    |
    // -------------------      V

    fetchProducts: (state : any, payload) => {
      const payloadProducts = payload.payload || [];
      state.products = payloadProducts;
      localStorage.setItem(applicationStateName, JSON.stringify({ ...state, externalCompanies: payloadProducts }));
      return state;
    },

    // -----------------------  |
    // -- INVOICES              |
    // -------------------      V

    fetchInvoices: (state : any, payload) => {
      const payloadInvoices = payload.payload || [];
      state.invoices = payloadInvoices;
      localStorage.setItem(applicationStateName, JSON.stringify({ ...state, invoices: payloadInvoices }));
      return state;
    },

    // -----------------------  |
    // -- INVOICE ITEMS         |
    // -------------------      V

    fetchInvoiceItems: (state : any, payload) => {
      const payloadInvoiceItems = payload.payload || [];
      state.invoiceItems = payloadInvoiceItems;
      localStorage.setItem(applicationStateName, JSON.stringify({ ...state, invoiceItems: payloadInvoiceItems }));
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
          projects: [],
          subProjects: [],
          externalCompanies: [],
          products: [],
          invoices: [],
          invoiceItems: [],
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

  fetchSubProjects,
  appendSubProject,
  popSubProject,
  selectProject,

  fetchExternalCompanies,

  fetchInvoices,
  fetchInvoiceItems,
  
  fetchProducts,
  
  clearData,
} = staticDataSlice.actions

export default staticDataSlice.reducer