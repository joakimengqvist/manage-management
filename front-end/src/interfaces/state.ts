/* eslint-disable @typescript-eslint/no-explicit-any */

import { SubProject } from "./subProject";

export interface StatePrivileges {
    [key: string]: {
        id: string
        name: string
    };
}

export interface StateUsers {
    [key: string]: {
        id: string
        first_name: string
        last_name: string
    };
}

export interface StateProjects {
    [key: string]: {
        id: string
        name: string
        sub_projects: Array<string>
    };
}

export interface StateSubProjects {
    [key: string]: {
        id: string
        name: string
        projects: Array<string>
    };
}

export interface StateExternalCompanies {
    [key: string]: {
        id: string
        company_name: string
    };
}

export interface StateProducts {
    [key: string]: {
        id: string
        name: string
        price: number
    };
}

export interface StateInvoices {
    [key: string]: {
        id: string
        company_id: string
        invoice_display_name: string
    }
}

export interface StateInvoiceItems {
    [key: string]: {
        id: string
        product_id: string
    }
}
export interface State {
    user: {
        id: string,
        authenticated: boolean
        firstName: string
        lastName: string
        email: string
        privileges: Array<string>
        projects: Array<string>
        settings: {
            dark_theme: boolean
            compact_ui: boolean
        }
    }
    application: {
        users: StateUsers
        privileges: StatePrivileges
        projects: StateProjects
        subProjects: Array<SubProject>
        externalCompanies: StateExternalCompanies
        products: StateProducts
        invoices: StateInvoices
        invoiceItems: StateInvoiceItems
    }
}