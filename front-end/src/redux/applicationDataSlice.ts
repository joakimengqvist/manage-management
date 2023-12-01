/* eslint-disable @typescript-eslint/no-explicit-any */
import { createSlice } from '@reduxjs/toolkit';
import {
  ExternalCompany,
  Invoice,
  InvoiceItem,
  Privilege,
  Product,
  Project,
  User,
  StateExternalCompanies,
  StateInvoiceItems,
  StateInvoices,
  StatePrivileges,
  StateProducts,
  StateProjects,
  StateUsers
} from '../interfaces';

export const applicationStateName = 'ReduxApplicationData';
const statename = 'application'

export const staticDataSlice = createSlice({
  name: statename,
  initialState: {
    selectedProject: {
      name: '',
      id: 0,
    },
    privileges: {},
    users: {},
    projects: {},
    subProjects: {},
    products: {},
    invoices: {},
    invoiceItems: {},
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

    fetchPrivileges: (state : any, payload) => {
      const privileges = payload.payload || [];
      const stateObject : StatePrivileges = {};
      privileges.forEach((privilege : Privilege) => {
        stateObject[privilege.id] = {
          id: privilege.id,
          name: privilege.name,
        }
      });
      state.privileges = stateObject;
      localStorage.setItem(applicationStateName, JSON.stringify({ ...state, privileges: stateObject }));
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

    fetchUsers: (state : any, payload) => {
      const users = payload.payload || [];
      const stateObject : StateUsers = {};
      users.forEach((user : User) => {
        stateObject[user.id] = {
          id: user.id,
          first_name: user.first_name,
          last_name: user.last_name,
        }
      });
      state.users = stateObject;
      localStorage.setItem(applicationStateName, JSON.stringify({ ...state, users: stateObject }));
      return state;
    },
    updateUserState: (state : any, payload) => {
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

    fetchProjects: (state : any, payload) => {
      const projects = payload.payload || [];
      const stateObject : StateProjects = {};
      projects.forEach((project : Project) => {
        stateObject[project.id] = {
          id: project.id,
          name: project.name,
          sub_projects: project.sub_projects,
        }
      });
      state.projects = stateObject;
      localStorage.setItem(applicationStateName, JSON.stringify({ ...state, projects: stateObject }));
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

    fetchSubProjects: (state : any, payload) => {
      const subProjects = payload.payload || [];
      state.subProjects = subProjects;
      localStorage.setItem(applicationStateName, JSON.stringify({ ...state, subProjects: subProjects }));
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

    fetchExternalCompanies: (state : any, payload) => {
      const externalCompanies = payload.payload || [];
      const stateObject : StateExternalCompanies = {};
      externalCompanies.forEach((externalCompany : ExternalCompany) => {
        stateObject[externalCompany.id] = {
          id: externalCompany.id,
          company_name: externalCompany.company_name,
        }
      });
      state.externalCompanies = stateObject;
      localStorage.setItem(applicationStateName, JSON.stringify({ ...state, externalCompanies: stateObject }));
      return state;
    },

    fetchProducts: (state : any, payload) => {
      const products = payload.payload || [];
      const stateObject : StateProducts = {};
      products.forEach((product : Product) => {
        stateObject[product.id] = {
          id: product.id,
          name: product.name,
          price: product.price,
        }
      });
      state.products = stateObject;
      localStorage.setItem(applicationStateName, JSON.stringify({ ...state, products: stateObject }));
      return state;
    },

    fetchInvoices: (state : any, payload) => {
      const invoices = payload.payload || [];
      const stateObject : StateInvoices = {};
      invoices.forEach((invoice : Invoice) => {
        stateObject[invoice.id] = {
          id: invoice.id,
          company_id: invoice.company_id,
        }
      });
      state.invoices = stateObject;
      localStorage.setItem(applicationStateName, JSON.stringify({ ...state, invoices: stateObject }));
      return state;
    },

    fetchInvoiceItems: (state : any, payload) => {
      const invoiceItems = payload.payload || [];
      const stateObject : StateInvoiceItems = {};
      invoiceItems.forEach((invoiceItem : InvoiceItem) => {
        stateObject[invoiceItem.id] = {
          id: invoiceItem.id,
          product_id: invoiceItem.product_id,
        }
      });
      state.invoiceItems = stateObject;
      localStorage.setItem(applicationStateName, JSON.stringify({ ...state, invoiceItems: stateObject }));
      return state;
    },

    clearData: state => {
      const updatedData = {
          ...state,
          selectedProject: {
            name: '',
            id: 0,
          },
          privileges: {},
          users: {},
          projects: {},
          subProjects: {},
          externalCompanies: {},
          products: {},
          invoices: {},
          invoiceItems: {},
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
  updateUserState,
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