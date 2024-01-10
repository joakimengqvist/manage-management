import { ENDPOINTS } from "./endpoints";

import { createExpense } from "./economics/expenses/create";
import { getExpenseById } from "./economics/expenses/getById";
import { updateExpense } from "./economics/expenses/update";
import { getAllExpenses } from "./economics/expenses/getAll";
import { getAllExpensesByProjectId } from "./economics/expenses/getAllByProjectId";

import { createIncome } from "./economics/incomes/create";
import { getIncomeById } from "./economics/incomes/getById";
import { updateIncome } from "./economics/incomes/update";
import { getAllIncomes } from "./economics/incomes/getAll";
import { getAllIncomesByProjectId } from "./economics/incomes/getAllByProjectId";

import { createExternalCompany } from "./externalCompanies/create";
import { getExternalCompanyById } from "./externalCompanies/getById";
import { updateExternalCompany } from "./externalCompanies/update";
import { getAllExternalCompanies } from "./externalCompanies/getAll";

import { createInvoice } from "./invoices/invoice/create";
import { getInvoiceById } from "./invoices/invoice/getById";
import { getInvoicesByIds } from "./invoices/invoice/getAllByIds";
import { updateInvoice } from "./invoices/invoice/update";
import { getAllInvoices } from "./invoices/invoice/getAll";
import { getAllInvoicesByProjectId } from "./invoices/invoice/getAllByProjectId";

import { createInvoiceItem } from "./invoices/invoiceItem/create";
import { getInvoiceItemById } from "./invoices/invoiceItem/getById";
import { updateInvoiceItem } from "./invoices/invoiceItem/update";
import { getAllInvoiceItems } from "./invoices/invoiceItem/getAll";

import { createExpenseNote } from "./notes/expense/create";
import { deleteExpenseNote } from "./notes/expense/delete";
import { getExpenseNoteById } from "./notes/expense/getById";
import { updateExpenseNote } from "./notes/expense/update";
import { getAllExpenseNotesByExpenseId } from "./notes/expense/getAllByExpenseId";
import { getAllExpenseNotesByUserId } from "./notes/expense/getAllByUserId";

import { createExternalCompanyNote } from "./notes/externalCompany/create";
import { deleteExternalCompanyNote } from "./notes/externalCompany/delete";
import { getExternalCompanyNoteById } from "./notes/externalCompany/getById";
import { updateExternalCompanyNote } from "./notes/externalCompany/update";
import { getAllExternalCompanyNotesByExternalCompanyId } from "./notes/externalCompany/getAllByExternalCompanyId";
import { getAllExternalCompanyNotesByUserId } from "./notes/externalCompany/getAllByUserId";
import { removeInvoiceFromCompany } from "./externalCompanies/specialActions/RemoveInvoiceFromCompany";
import { addInvoiceToCompany } from "./externalCompanies/specialActions/AddInvoiceToCompany";

import { createProductNote } from "./notes/product/create";
import { deleteProductNote } from "./notes/product/delete";
import { getProductNoteById } from "./notes/product/getById";
import { updateProductNote } from "./notes/product/update";
import { getAllProductNotesByProductId } from "./notes/product/getAllByProductId";
import { getAllProductNotesByUserId } from "./notes/product/getAllByUserId";

import { createIncomeNote } from "./notes/income/create";
import { deleteIncomeNote } from "./notes/income/delete";
import { getIncomeNoteById } from "./notes/income/getById";
import { updateIncomeNote } from "./notes/income/update";
import { getAllIncomeNotesByIncomeId } from "./notes/income/getAllByIncomeId";
import { getAllIncomeNotesByUserId } from "./notes/income/getAllByUserId";

import { createProjectNote } from "./notes/project/create";
import { deleteProjectNote } from "./notes/project/delete";
import { getProjectNoteById } from "./notes/project/getById";
import { updateProjectNote } from "./notes/project/update";
import { getAllProjectNotesByProjectId } from "./notes/project/getAllByProjectd";  
import { getAllProjectNotesByUserId } from "./notes/project/getAllByUserId";

import { createSubProjectNote } from "./notes/subProject/create";
import { deleteSubProjectNote } from "./notes/subProject/delete";
import { getSubProjectNoteById } from "./notes/subProject/getById";
import { updateSubProjectNote } from "./notes/subProject/update";
import { getAllSubProjectNotesBySubProjectId } from "./notes/subProject/getAllBySubProjectId";
import { getAllSubProjectNotesByUserId } from "./notes/subProject/getAllByUserId";

import { createInvoiceNote } from "./notes/invoice/create";
import { deleteInvoiceNote } from "./notes/invoice/delete";
import { getInvoiceNoteById } from "./notes/invoice/getById";
import { updateInvoiceNote } from "./notes/invoice/update";
import { getAllInvoiceNotesByInvoiceId } from "./notes/invoice/getAllByInvoiceId";  
import { getAllInvoiceNotesByUserId } from "./notes/invoice/getAllByUserId";

import { createInvoiceItemNote } from "./notes/invoiceItem/create";
import { deleteInvoiceItemNote } from "./notes/invoiceItem/delete";
import { getInvoiceItemNoteById } from "./notes/invoiceItem/getById";
import { updateInvoiceItemNote } from "./notes/invoiceItem/update";
import { getAllInvoiceItemNotesByInvoiceItemId } from "./notes/invoiceItem/getAllByInvoiceItemId";  
import { getAllInvoiceItemNotesByUserId } from "./notes/invoiceItem/getAllByUserId";

import { createPrivilege } from "./privileges/create";
import { getPrivilegeById } from "./privileges/getById";
import { updatePrivilege } from "./privileges/update";
import { getAllPrivileges } from "./privileges/getAll";
import { deletePrivilege } from "./privileges/delete";

import { createProduct } from "./products/create";
import { getProductById } from "./products/getById";
import { getProductsByIds } from "./products/getAllByIds";
import { updateProduct } from "./products/update";
import { getAllProducts } from "./products/getAll";

import { createProject } from "./projects/create";
import { getProjectById } from "./projects/getById";
import { updateProject } from "./projects/update";
import { getAllProjects } from "./projects/getAll";
import { getProjectsByIds } from "./projects/getAllByIds";

import { createSubProject } from "./subProjects/create";
import { getSubProjectById } from "./subProjects/getById";
import { updateSubProject } from "./subProjects/update";
import { getAllSubProjects } from "./subProjects/getAll";
import { getSubProjectsByIds } from "./subProjects/getAllByIds";
import { deleteSubProject } from "./subProjects/delete";
import { addProjectsSubProjectConnection } from "./subProjects/specialActions/addProjectsSubProjectConnection";
import { removeProjectsSubProjectConnection } from "./subProjects/specialActions/removeProjectsSubProjectConnection";
import { addSubProjectsProjectConnection } from "./subProjects/specialActions/addSubProjectsProjectConnection";
import { removeSubProjectsProjectConnection } from "./subProjects/specialActions/removeSubProjectsProjectsConnection";

import { createUser } from "./users/create";
import { getUserById } from "./users/getById";
import { updateUser } from "./users/update";
import { getAllUsers } from "./users/getAll";
import { deleteUser } from "./users/delete";
import { loginAuthenticate } from "./users/loginAuthenticate";
import { getUserSettingsByUserId } from "./users/userSettings/GetUserSettingsByUserId";
import { updateUserSettings } from "./users/userSettings/updateUserSettings";

export {
    ENDPOINTS,
    
    createExpense,
    getExpenseById,
    updateExpense,
    getAllExpenses,
    getAllExpensesByProjectId,

    createIncome,
    getIncomeById,
    updateIncome,
    getAllIncomes,
    getAllIncomesByProjectId,
    
    createExternalCompany,
    getExternalCompanyById,
    updateExternalCompany,
    getAllExternalCompanies,
    
    createInvoice,
    getInvoiceById,
    getInvoicesByIds,
    updateInvoice,
    getAllInvoices,
    getAllInvoicesByProjectId,
    
    createInvoiceItem,
    getInvoiceItemById,
    updateInvoiceItem,
    getAllInvoiceItems,
    
    createExpenseNote,
    deleteExpenseNote,
    getExpenseNoteById,
    updateExpenseNote,
    getAllExpenseNotesByExpenseId,
    getAllExpenseNotesByUserId,
        
    createExternalCompanyNote,
    deleteExternalCompanyNote,
    getExternalCompanyNoteById,
    updateExternalCompanyNote,
    getAllExternalCompanyNotesByExternalCompanyId,
    getAllExternalCompanyNotesByUserId,
    removeInvoiceFromCompany,
    addInvoiceToCompany,
     
    createIncomeNote,
    deleteIncomeNote,
    getIncomeNoteById, 
    updateIncomeNote,
    getAllIncomeNotesByIncomeId,
    getAllIncomeNotesByUserId,

    createProjectNote,
    deleteProjectNote,
    getProjectNoteById,
    updateProjectNote,
    getAllProjectNotesByProjectId,
    getAllProjectNotesByUserId,

    createProductNote,
    deleteProductNote,
    getProductNoteById,
    updateProductNote,
    getAllProductNotesByProductId,
    getAllProductNotesByUserId,
    
    createSubProjectNote,
    deleteSubProjectNote,
    getSubProjectNoteById,
    updateSubProjectNote,
    getAllSubProjectNotesBySubProjectId,
    getAllSubProjectNotesByUserId,

    createInvoiceNote,
    deleteInvoiceNote,
    getInvoiceNoteById,
    updateInvoiceNote,
    getAllInvoiceNotesByInvoiceId,
    getAllInvoiceNotesByUserId,

    createInvoiceItemNote,
    deleteInvoiceItemNote,
    getInvoiceItemNoteById,
    updateInvoiceItemNote,
    getAllInvoiceItemNotesByInvoiceItemId,
    getAllInvoiceItemNotesByUserId,
    
    createPrivilege,
    getPrivilegeById,
    updatePrivilege,
    getAllPrivileges,
    deletePrivilege,
    
    createProduct,
    getProductById,
    getProductsByIds,
    updateProduct,
    getAllProducts,
    
    createProject,
    getProjectById,
    updateProject,
    getAllProjects,
    getProjectsByIds,
    
    createSubProject,
    getSubProjectById,
    updateSubProject,
    getAllSubProjects,
    getSubProjectsByIds,
    deleteSubProject,
    addProjectsSubProjectConnection,
    removeProjectsSubProjectConnection,
    addSubProjectsProjectConnection,
    removeSubProjectsProjectConnection,
    
    createUser,
    getUserById,
    updateUser,
    getAllUsers,
    deleteUser,
    loginAuthenticate,
    getUserSettingsByUserId,
    updateUserSettings
}
