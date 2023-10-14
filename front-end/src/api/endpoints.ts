export enum ENDPOINTS {

    // USERS
    Authenticate = "http://localhost:8080/auth/authenticate",
    CreateUser = "http://localhost:8080/auth/create-user",
    UpdateUser = "http://localhost:8080/auth/update-user",
    DeleteUser = "http://localhost:8080/auth/delete-user",
    GetUserById = "http://localhost:8080/auth/get-user-by-id",
    GetAllUsers = "http://localhost:8080/auth/get-all-users",

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
    GetAllProjects = "http://localhost:8080/project/get-all-projects",

    // PROJECT NOTES
    CreateProjectNote = "http://localhost:8080/notes/create-project-note",
    GetProjectNoteById = "http://localhost:8080/notes/get-project-note-by-id",
    UpdateProjectNote = "http://localhost:8080/notes/update-project-note",
    GetAllProjectNotesByProjectId = "http://localhost:8080/notes/get-all-project-notes-by-project-id",
    GetAllProjectNotesByUserId = "http://localhost:8080/notes/get-all-project-notes-by-user-id",
    DeleteProjectNote = "http://localhost:8080/notes/delete-project",

    // PROJECT INCOMES
    CreateProjectIncome = "http://localhost:8080/economics/create-project-income",
    GetProjectIncomeById = "http://localhost:8080/economics/get-project-income-by-id",
    UpdateProjectIncome = "http://localhost:8080/economics/update-project-income",
    GetAllProjectIncomes = "http://localhost:8080/economics/get-all-project-incomes",
    GetAllProjectIncomesByProjectId = "http://localhost:8080/economics/get-all-project-incomes-by-project-id",
    GetAllProjectIncomesByUserId = "http://localhost:8080/economics/get-all-project-incomes-by-user-id",
    DeleteProjectIncome = "http://localhost:8080/economics/delete-project-income",

    // INCOME NOTES
    CreateIncomeNote = "http://localhost:8080/notes/create-income-note",
    GetIncomeNoteById = "http://localhost:8080/notes/get-income-note-by-id",
    UpdateIncomeNote = "http://localhost:8080/notes/update-income-note",
    GetAllIncomeNotesByIncomeId = "http://localhost:8080/notes/get-all-income-notes-by-income-id",
    GetAllIncomeNotesByUserId = "http://localhost:8080/notes/get-all-income-notes-by-user-id",
    DeleteIncomeNote = "http://localhost:8080/notes/delete-income-note",

    // PROJECT EXPENSES
    CreateProjectExpense = "http://localhost:8080/economics/create-project-expense",
    GetProjectExpenseById = "http://localhost:8080/economics/get-project-expense-by-id",
    UpdateProjectExpense = "http://localhost:8080/economics/update-project-expense",
    getAllProectExpenses = "http://localhost:8080/economics/get-all-project-expenses",
    GetAllProjectExpensesByProjectId = "http://localhost:8080/economics/get-all-project-expenses-by-project-id",
    GetAllProjectExpensesByUserId = "http://localhost:8080/economics/get-all-project-expenses-by-user-id",
    DeleteProjectExpense = "http://localhost:8080/economics/delete-project-expense",

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

    // EMAIL
    sendEmail = "http://localhost:8080/email/send-email",
}
