/* eslint-disable @typescript-eslint/no-explicit-any */

import { ExternalCompany } from "./externalCompany"
import { Invoice, InvoiceItem } from "./invoice"
import { Privilege } from "./privilege"
import { Product } from "./product"
import { Project } from "./project"
import { SubProject } from "./subProject"
import { User } from "./user"

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
        privileges: Array<Privilege>
        projects: Array<Project>
        subProjects: Array<SubProject>
        users: Array<User>
        externalCompanies: Array<ExternalCompany>
        products: Array<Product>
        invoices: Array<Invoice>
        invoiceItems: Array<InvoiceItem>
    }
}