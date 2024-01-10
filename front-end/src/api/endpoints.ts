// 192.168.49.2:30002
/*
console.log('process.env', process?.env);
const domain = process?.env?.DEVELOPMENT_DOCKER_COMPOSE ? 
"http://localhost:8080" : 
`http://${process.env.BROKER_SERVICE_SERVICE_HOST}`;
*/


const domain = "http://localhost:8080";



export const ENDPOINTS = {

    // USERS
    Authenticate: `${domain}/auth/authenticate`,
    CreateUser: `${domain}/auth/create-user`,
    UpdateUser: `${domain}/auth/update-user`,
    DeleteUser: `${domain}/auth/delete-user`,
    GetUserById: `${domain}/auth/get-user-by-id`,
    GetAllUsers: `${domain}/auth/get-all-users`,

    // USER SETTINGS
    GetUserSettingsByUserId: `${domain}/auth/get-user-settings-by-user-id`,
    UpdateUserSettings: `${domain}/auth/update-user-settings`,

    // PRIVILEGES
    CreatePrivilege: `${domain}/auth/create-privilege`,
    UpdatePrivilege: `${domain}/auth/update-privilege`,
    DeletePrivilege: `${domain}/auth/delete-privilege`,
    GetPrivilegeById: `${domain}/auth/get-privilege-by-id`,
    GetAllPrivileges: `${domain}/auth/get-all-privileges`,

    // PROJECTS
    CreateProject: `${domain}/project/create-project`,
    UpdateProject: `${domain}/project/update-project`,
    DeleteProject: `${domain}/project/delete-project`,
    GetProjectById: `${domain}/project/get-project-by-id`,
    GetProjectsByIds: `${domain}/project/get-projects-by-ids`,
    GetAllProjects: `${domain}/project/get-all-projects`,

    // PROJECT NOTES
    CreateProjectNote: `${domain}/notes/create-project-note`,
    GetProjectNoteById: `${domain}/notes/get-project-note-by-id`,
    UpdateProjectNote: `${domain}/notes/update-project-note`,
    GetAllProjectNotesByProjectId: `${domain}/notes/get-all-project-notes-by-project-id`,
    GetAllProjectNotesByUserId: `${domain}/notes/get-all-project-notes-by-user-id`,
    DeleteProjectNote: `${domain}/notes/delete-project-note`,

    // SUBPROJECTS
    CreateSubProject: `${domain}/project/create-sub-project`,
    UpdateSubProject: `${domain}/project/update-sub-project`,
    DeleteSubProject: `${domain}/project/delete-sub-project`,
    GetSubProjectById: `${domain}/project/get-sub-project-by-id`,
    GetSubProjectsByIds: `${domain}/project/get-sub-projects-by-ids`,
    GetAllSubProjects: `${domain}/project/get-all-sub-projects`,

    AddProjectsSubProjectConnection: `${domain}/project/add-projects-sub-project-connection`,
    RemoveProjectsSubProjectConnection: `${domain}/project/delete-projects-sub-project-connection`,
    AddSubProjectsProjectConnection: `${domain}/project/add-sub-projects-project-connection`,
    RemoveSubProjectsProjectConnection: `${domain}/project/delete-sub-projects-project-connection`,

    // SUB PROJECT NOTES
    CreateSubProjectNote: `${domain}/notes/create-sub-project-note`,
    GetSubProjectNoteById: `${domain}/notes/get-sub-project-note-by-id`,
    UpdateSubProjectNote: `${domain}/notes/update-sub-project-note`,
    GetAllSubProjectNotesBySubProjectId: `${domain}/notes/get-all-sub-project-notes-by-sub-project-id`,
    GetAllSubProjectNotesByUserId: `${domain}/notes/get-all-sub-project-notes-by-user-id`,
    DeleteSubProjectNote: `${domain}/notes/delete-sub-project-note`,

    // INCOMES
    CreateIncome: `${domain}/economics/create-income`,
    GetIncomeById: `${domain}/economics/get-income-by-id`,
    UpdateIncome: `${domain}/economics/update-income`,
    GetAllIncomes: `${domain}/economics/get-all-incomes`,
    GetAllIncomesByProjectId: `${domain}/economics/get-all-incomes-by-project-id`,
    GetAllIncomesByUserId: `${domain}/economics/get-all-incomes-by-user-id`,
    DeleteIncome: `${domain}/economics/delete-income`,

    // INCOME NOTES
    CreateIncomeNote: `${domain}/notes/create-income-note`,
    GetIncomeNoteById: `${domain}/notes/get-income-note-by-id`,
    UpdateIncomeNote: `${domain}/notes/update-income-note`,
    GetAllIncomeNotesByIncomeId: `${domain}/notes/get-all-income-notes-by-income-id`,
    GetAllIncomeNotesByUserId: `${domain}/notes/get-all-income-notes-by-user-id`,
    DeleteIncomeNote: `${domain}/notes/delete-income-note`,

    // EXPENSES
    CreateExpense: `${domain}/economics/create-expense`,
    GetExpenseById: `${domain}/economics/get-expense-by-id`,
    UpdateExpense: `${domain}/economics/update-expense`,
    getAllExpenses: `${domain}/economics/get-all-expenses`,
    GetAllExpensesByProjectId: `${domain}/economics/get-all-expenses-by-project-id`,
    GetAllExpensesByUserId: `${domain}/economics/get-all-expenses-by-user-id`,
    DeleteExpense: `${domain}/economics/delete-expense`,

    // EXPENSE NOTES
    CreateExpenseNote: `${domain}/notes/create-expense-note`,
    GetExpenseNoteById: `${domain}/notes/get-expense-note-by-id`,
    UpdateExpenseNote: `${domain}/notes/update-expense-note`,
    GetAllExpenseNotesByExpenseId: `${domain}/notes/get-all-expense-notes-by-expense-id`,
    GetAllExpenseNotesByUserId: `${domain}/notes/get-all-expense-notes-by-user-id`,
    DeleteExpenseNote: `${domain}/notes/delete-expense-note`,

    // EXTERNAL COMPANIES
    CreateExternalCompany: `${domain}/external-company/create-external-company`,
    GetExternalCompanyById: `${domain}/external-company/get-external-company-by-id`,
    GetAllExternalCompanies: `${domain}/external-company/get-all-external-companies`,
    UpdateExternalCompany: `${domain}/external-company/update-external-company`,
    DeleteExternalCompany: `${domain}/external-company/delete-external-company`,
    RemoveInvoiceFromCompany: `${domain}/external-company/remove-invoice`,
    addInvoiceFromCompany: `${domain}/external-company/append-invoice`,

    // EXTERNAL COMPANY NOTES
    CreateExternalCompanyNote: `${domain}/notes/create-external-company-note`,
    GetExternalCompanyNoteById: `${domain}/notes/get-external-company-note-by-id`,
    UpdateExternalCompanyNote: `${domain}/notes/update-external-company-note`,
    GetAllExternalCompanyNotesByExternalCompanyId: `${domain}/notes/get-all-external-company-notes-by-external-company-id`,
    GetAllExternalCompanyNotesByUserId: `${domain}/notes/get-all-external-company-notes-by-user-id`,
    DeleteExternalCompanyNote: `${domain}/notes/delete-external-company-note`,

    // PRODUCTS
    CreateProduct: `${domain}/product/create-product`,
    GetProductById: `${domain}/product/get-product-by-id`,
    GetAllProducts: `${domain}/product/get-all-products`,
    UpdateProduct: `${domain}/product/update-product`,
    getProductsByIds: `${domain}/product/get-products-by-ids`,

    // PRODUCT NOTES
    CreateProductNote: `${domain}/notes/create-product-note`,
    GetProductNoteById: `${domain}/notes/get-product-note-by-id`,
    UpdateProductNote: `${domain}/notes/update-product-note`,
    GetAllProductNotesByProductId: `${domain}/notes/get-all-product-notes-by-product-id`,
    GetAllProductNotesByUserId: `${domain}/notes/get-all-product-notes-by-user-id`,
    DeleteProductNote: `${domain}/notes/delete-product-note`,

    // INVOICES
    CreateInvoice: `${domain}/invoice/create-invoice`,
    GetInvoiceById: `${domain}/invoice/get-invoice-by-id`,
    GetAllInvoices: `${domain}/invoice/get-all-invoices`,
    getAllInvoicesByIds: `${domain}/invoice/get-all-invoices-by-ids`,
    GetAllInvoicesByProjectId: `${domain}/invoice/get-all-invoices-by-project-id`,
    GetAllInvoicesBySubProjectId: `${domain}/invoice/get-all-invoices-by-sub-project-id`,
    UpdateInvoice: `${domain}/invoice/update-invoice`,
    DeleteInvoice: `${domain}/invoice/delete-invoice`,

    // INVOICE NOTES
    CreateInvoiceNote: `${domain}/notes/create-invoice-note`,
    GetInvoiceNoteById: `${domain}/notes/get-invoice-note-by-id`,
    UpdateInvoiceNote: `${domain}/notes/update-invoice-note`,
    GetAllInvoiceNotesByInvoiceId: `${domain}/notes/get-all-invoice-notes-by-invoice-id`,
    GetAllInvoiceNotesByUserId: `${domain}/notes/get-all-invoice-notes-by-user-id`,
    DeleteInvoiceNote: `${domain}/notes/delete-invoice-note`,

    // INVOICE ITEMS
    CreateInvoiceItem: `${domain}/invoice/create-invoice-item`,
    GetInvoiceItemById: `${domain}/invoice/get-invoice-item-by-id`,
    GetAllInvoiceItems: `${domain}/invoice/get-all-invoice-items`,
    getAllInvoiceItemsByIds: `${domain}/invoice/get-all-invoice-items-by-ids`,
    UpdateInvoiceItem: `${domain}/invoice/update-invoice-item`,
    DeleteInvoiceItem: `${domain}/invoice/delete-invoice-item`,
    GetAllInvoiceItemsByInvoiceId: `${domain}/invoice/get-all-invoice-items-by-invoice-id`,

    // INVOICE NOTES
    CreateInvoiceItemNote: `${domain}/notes/create-invoice-item-note`,
    GetInvoiceItemNoteById: `${domain}/notes/get-invoice-item-note-by-id`,
    UpdateInvoiceItemNote: `${domain}/notes/update-invoice-item-note`,
    GetAllInvoiceItemNotesByInvoiceItemId: `${domain}/notes/get-all-invoice-item-notes-by-invoice-item-id`,
    GetAllInvoiceItemNotesByUserId: `${domain}/notes/get-all-invoice-item-notes-by-user-id`,
    DeleteInvoiceItemNote: `${domain}/notes/delete-invoice-item-note`,
} as const;