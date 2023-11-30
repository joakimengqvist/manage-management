export enum ENDPOINTS {

    // USERS
    Authenticate = "http://localhost:8080/auth/authenticate",
    CreateUser = "http://localhost:8080/auth/create-user",
    UpdateUser = "http://localhost:8080/auth/update-user",
    DeleteUser = "http://localhost:8080/auth/delete-user",
    GetUserById = "http://localhost:8080/auth/get-user-by-id",
    GetAllUsers = "http://localhost:8080/auth/get-all-users",

    // USER SETTINGS
    GetUserSettingsByUserId = "http://localhost:8080/auth/get-user-settings-by-user-id",
    UpdateUserSettings = "http://localhost:8080/auth/update-user-settings",

    // PRIVILEGES
    CreatePrivilege = "http://localhost:8080/auth/create-privilege",
    UpdatePrivilege = "http://localhost:8080/auth/update-privilege",
    DeletePrivilege = "http://localhost:8080/auth/delete-privilege",
    GetPrivilegeById = "http://localhost:8080/auth/get-privilege-by-id",
    GetAllPrivileges = "http://localhost:8080/auth/get-all-privileges",

    // PROJECTS
    CreateProject = "http://localhost:8080/project/create-project",
    UpdateProject = "http://localhost:8080/project/update-project",
    DeleteProject = "http://localhost:8080/project/delete-project",
    GetProjectById = "http://localhost:8080/project/get-project-by-id",
    GetProjectsByIds = "http://localhost:8080/project/get-projects-by-ids",
    GetAllProjects = "http://localhost:8080/project/get-all-projects",


    // PROJECT NOTES
    CreateProjectNote = "http://localhost:8080/notes/create-project-note",
    GetProjectNoteById = "http://localhost:8080/notes/get-project-note-by-id",
    UpdateProjectNote = "http://localhost:8080/notes/update-project-note",
    GetAllProjectNotesByProjectId = "http://localhost:8080/notes/get-all-project-notes-by-project-id",
    GetAllProjectNotesByUserId = "http://localhost:8080/notes/get-all-project-notes-by-user-id",
    DeleteProjectNote = "http://localhost:8080/notes/delete-project-note",

    // SUBPROJECTS
    CreateSubProject = "http://localhost:8080/project/create-sub-project",
    UpdateSubProject = "http://localhost:8080/project/update-sub-project",
    DeleteSubProject = "http://localhost:8080/project/delete-sub-project",
    GetSubProjectById = "http://localhost:8080/project/get-sub-project-by-id",
    GetSubProjectsByIds = "http://localhost:8080/project/get-sub-projects-by-ids",
    GetAllSubProjects = "http://localhost:8080/project/get-all-sub-projects",

    AddProjectsSubProjectConnection = "http://localhost:8080/project/add-projects-sub-project-connection",
    RemoveProjectsSubProjectConnection = "http://localhost:8080/project/delete-projects-sub-project-connection",
    AddSubProjectsProjectConnection = "http://localhost:8080/project/add-sub-projects-project-connection",
    RemoveSubProjectsProjectConnection = "http://localhost:8080/project/delete-sub-projects-project-connection",

    // PROJECT NOTES
    CreateSubProjectNote = "http://localhost:8080/notes/create-sub-project-note",
    GetSubProjectNoteById = "http://localhost:8080/notes/get-sub-project-note-by-id",
    UpdateSubProjectNote = "http://localhost:8080/notes/update-sub-project-note",
    GetAllSubProjectNotesBySubProjectId = "http://localhost:8080/notes/get-all-sub-project-notes-by-sub-project-id",
    GetAllSubProjectNotesByUserId = "http://localhost:8080/notes/get-all-sub-project-notes-by-user-id",
    DeleteSubProjectNote = "http://localhost:8080/notes/delete-sub-project-note",

    // INCOMES
    CreateIncome = "http://localhost:8080/economics/create-income",
    GetIncomeById = "http://localhost:8080/economics/get-income-by-id",
    UpdateIncome = "http://localhost:8080/economics/update-income",
    GetAllIncomes = "http://localhost:8080/economics/get-all-incomes",
    GetAllIncomesByProjectId = "http://localhost:8080/economics/get-all-incomes-by-project-id",
    GetAllIncomesByUserId = "http://localhost:8080/economics/get-all-incomes-by-user-id",
    DeleteIncome = "http://localhost:8080/economics/delete-income",

    // INCOME NOTES
    CreateIncomeNote = "http://localhost:8080/notes/create-income-note",
    GetIncomeNoteById = "http://localhost:8080/notes/get-income-note-by-id",
    UpdateIncomeNote = "http://localhost:8080/notes/update-income-note",
    GetAllIncomeNotesByIncomeId = "http://localhost:8080/notes/get-all-income-notes-by-income-id",
    GetAllIncomeNotesByUserId = "http://localhost:8080/notes/get-all-income-notes-by-user-id",
    DeleteIncomeNote = "http://localhost:8080/notes/delete-income-note",

    // EXPENSES
    CreateExpense = "http://localhost:8080/economics/create-expense",
    GetExpenseById = "http://localhost:8080/economics/get-expense-by-id",
    UpdateExpense = "http://localhost:8080/economics/update-expense",
    getAllExpenses = "http://localhost:8080/economics/get-all-expenses",
    GetAllExpensesByProjectId = "http://localhost:8080/economics/get-all-expenses-by-project-id",
    GetAllExpensesByUserId = "http://localhost:8080/economics/get-all-expenses-by-user-id",
    DeleteExpense = "http://localhost:8080/economics/delete-expense",

    // EXPENSE NOTES
    CreateExpenseNote = "http://localhost:8080/notes/create-expense-note",
    GetExpenseNoteById = "http://localhost:8080/notes/get-expense-note-by-id",
    UpdateExpenseNote = "http://localhost:8080/notes/update-expense-note",
    GetAllExpenseNotesByExpenseId = "http://localhost:8080/notes/get-all-expense-notes-by-expense-id",
    GetAllExpenseNotesByUserId = "http://localhost:8080/notes/get-all-expense-notes-by-user-id",
    DeleteExpenseNote = "http://localhost:8080/notes/delete-expense-note",

    // EXTERNAL COMPANIES
    CreateExternalCompany = "http://localhost:8080/external-company/create-external-company",
    GetExternalCompanyById = "http://localhost:8080/external-company/get-external-company-by-id",
    GetAllExternalCompanies = "http://localhost:8080/external-company/get-all-external-companies",
    UpdateExternalCompany = "http://localhost:8080/external-company/update-external-company",
    DeleteExternalCompany = "http://localhost:8080/external-company/delete-external-company",

    // EXTERNAL COMPANY NOTES
    CreateExternalCompanyNote = "http://localhost:8080/notes/create-external-company-note",
    GetExternalCompanyNoteById = "http://localhost:8080/notes/get-external-company-note-by-id",
    UpdateExternalCompanyNote = "http://localhost:8080/notes/update-external-company-note",
    GetAllExternalCompanyNotesByExternalCompanyId = "http://localhost:8080/notes/get-all-external-company-notes-by-external-company-id",
    GetAllExternalCompanyNotesByUserId = "http://localhost:8080/notes/get-all-external-company-notes-by-user-id",
    DeleteExternalCompanyNote = "http://localhost:8080/notes/delete-external-company-note",

    // PRODUCTS
    CreateProduct = "http://localhost:8080/product/create-product",
    GetProductById = "http://localhost:8080/product/get-product-by-id",
    GetAllProducts = "http://localhost:8080/product/get-all-products",
    UpdateProduct = "http://localhost:8080/product/update-product",
    getProductsByIds = "http://localhost:8080/product/get-products-by-ids",

    // INVOICES
    CreateInvoice = "http://localhost:8080/invoice/create-invoice",
    GetInvoiceById = "http://localhost:8080/invoice/get-invoice-by-id",
    GetAllInvoices = "http://localhost:8080/invoice/get-all-invoices",
    GetAllInvoicesByProjectId = "http://localhost:8080/invoice/get-all-invoices-by-project-id",
    GetAllInvoicesBySubProjectId = "http://localhost:8080/invoice/get-all-invoices-by-sub-project-id",
    UpdateInvoice = "http://localhost:8080/invoice/update-invoice",
    DeleteInvoice = "http://localhost:8080/invoice/delete-invoice",

    // INVOICE ITEMS
    CreateInvoiceItem = "http://localhost:8080/invoice/create-invoice-item",
    GetInvoiceItemById = "http://localhost:8080/invoice/get-invoice-item-by-id",
    GetAllInvoiceItems = "http://localhost:8080/invoice/get-all-invoice-items",
    getAllInvoiceItemsByIds = "http://localhost:8080/invoice/get-all-invoice-items-by-ids",
    UpdateInvoiceItem = "http://localhost:8080/invoice/update-invoice-item",
    DeleteInvoiceItem = "http://localhost:8080/invoice/delete-invoice-item",
    GetAllInvoiceItemsByInvoiceId = "http://localhost:8080/invoice/get-all-invoice-items-by-invoice-id",


    // EMAIL
    sendEmail = "http://localhost:8080/email/send-email",
}
