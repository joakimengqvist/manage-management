/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable react-hooks/exhaustive-deps */
import { useState, useEffect } from 'react';
import { Route, Routes, useNavigate } from 'react-router-dom';
import { useDispatch } from 'react-redux';
import './App.css';
import HeaderMenu from './components/header/HeaderMenu';
import Navigation from './components/navigation/Navigation';
import Login from './routes/Login';
import UsersDetails from './routes/user/Users';
import User from './routes/user/User';
import Projects from './routes/project/Projects';
import Project from './routes/project/Project';
import Privileges from './routes/privilege/Privileges';
import Privilege from './routes/privilege/Privilege';
import { ConfigProvider, Layout, theme } from 'antd';
import { getAllPrivileges } from './api/privileges/getAll';
import { fetchExternalCompanies, fetchInvoiceItems, fetchInvoices, fetchPrivileges, fetchProducts, fetchProjects, fetchSubProjects, fetchUsers, initiateApplicationData } from './redux/applicationDataSlice';
import { fetchUserSettings, initiateUser } from './redux/userDataSlice'
import { getAllUsers } from './api/users/getAll';
import { getAllProjects } from './api/projects/getAll';
import { hasPrivilege } from './helpers/hasPrivileges';
import { PRIVILEGES } from './enums/privileges';
import MyPage from './routes/myPage';
import ExpenseCreate from './routes/expense/ExpenseCreate';
import Expense from './routes/expense/Expense';
import Expenses from './routes/expense/Expenses';
import ServiceOverviewPage from './routes/ServiceOverview';
import Incomes from './routes/income/Incomes';
// import IncomeCreate from './routes/income/IncomeCreate';
import Income from './routes/income/Income';
import ExternalCompanies from './routes/externalCompany/ExternalCompanies';
import ExternalCompanyCreate from './routes/externalCompany/ExternalCompanyCreate';
import ExternalCompany from './routes/externalCompany/ExternalCompany';
import { getAllExternalCompanies } from './api/externalCompanies/getAll';
import SubProjects from './routes/subProject/SubProjects';
import SubProjectCreate from './routes/subProject/SubProjectCreate';
import { getAllSubProjects } from './api/subProjects/getAll';
import ProjectCreate from './routes/project/ProjectCreate';
import SubProject from './routes/subProject/SubProject';
import Testing from './routes/Testing';
import { getAllProducts } from './api/products/getAll';
import Products from './routes/product/Products';
import CreateInvoice from './routes/invoice/InvoiceCreate';
import Invoice from './routes/invoice/Invoice';
import Invoices from './routes/invoice/Invoices';
import InvoiceItem from './routes/invoice/invoiceItem';
import InvoiceItems from './routes/invoice/invoiceItems';
import { getAllInvoiceItems } from './api/invoices/invoiceItem/getAll';
import { getAllInvoices } from './api/invoices/invoice/getAll';
import { getUserSettingsByUserId } from './api/users/userSettings/GetUserSettingsByUserId';
import Documentation from './routes/Documentation';
import { useGetCompactUIEnabled, useGetDarkThemeEnabled, useGetLoggedInUserId } from './hooks';
import { useGetAuthenticated } from './hooks/useGetAuthenticated';
import { useGetLoggedInUserPrivileges } from './hooks/useGetLoggedInUserPrivileges';

const { Sider, Content } = Layout;

const App = () => {
  const dispatch = useDispatch();
  const navigate = useNavigate();  
  const authenticated = useGetAuthenticated();
  const loggedInUserId = useGetLoggedInUserId();
  const loggedInUserPrivileges = useGetLoggedInUserPrivileges();
  const darkTheme = useGetDarkThemeEnabled();
  const compactUI = useGetCompactUIEnabled();
  const [collapsed, setCollapsed] = useState(false);

  if (!authenticated) {
    navigate('/login')
  }

  useEffect(() => {
      dispatch(initiateUser());
      dispatch(initiateApplicationData());

    if (loggedInUserId && authenticated) {
      getUserSettingsByUserId(loggedInUserId, loggedInUserId).then(response => dispatch(fetchUserSettings(response.data))).catch(() => {})

      if (hasPrivilege(loggedInUserPrivileges, PRIVILEGES.privilege_read)) {
        getAllPrivileges(loggedInUserId).then(response => dispatch(fetchPrivileges(response.data))).catch(() => {})
      }
      if (hasPrivilege(loggedInUserPrivileges, PRIVILEGES.user_read)) {
        getAllUsers(loggedInUserId).then(response => dispatch(fetchUsers(response.data))).catch(() => {})
      }
      if (hasPrivilege(loggedInUserPrivileges, PRIVILEGES.project_read)) {
        getAllProjects(loggedInUserId).then(response => dispatch(fetchProjects(response.data))).catch(() => {})
      }
      if (hasPrivilege(loggedInUserPrivileges, PRIVILEGES.sub_project_read)) {
        getAllSubProjects(loggedInUserId).then(response => dispatch(fetchSubProjects(response.data))).catch(() => {})
      }
      if (hasPrivilege(loggedInUserPrivileges, PRIVILEGES.external_company_read)) {
        getAllExternalCompanies(loggedInUserId).then(response => dispatch(fetchExternalCompanies(response.data))).catch(() => {})
      }
      if (hasPrivilege(loggedInUserPrivileges, PRIVILEGES.product_read)) {
        getAllProducts(loggedInUserId).then(response => dispatch(fetchProducts(response.data))).catch(() => {})
      }
      if (hasPrivilege(loggedInUserPrivileges, PRIVILEGES.invoice_read)) {
        getAllInvoices(loggedInUserId).then(response => dispatch(fetchInvoices(response.data))).catch(() => {})
      }
      if (hasPrivilege(loggedInUserPrivileges, PRIVILEGES.invoice_read)) {
        getAllInvoiceItems(loggedInUserId).then(response => dispatch(fetchInvoiceItems(response.data))).catch(() => {})
      }
    }
}, [loggedInUserId, authenticated])

const ThemePicker = () => {
  if (compactUI) {
    if (darkTheme) {
      return [theme.darkAlgorithm, theme.compactAlgorithm]
    } else {
      return [theme.compactAlgorithm]
    }
  }

  return [darkTheme ? theme.darkAlgorithm : theme.defaultAlgorithm]
}

const isDocPage = () => {
  if (window.location.pathname.includes('documentation')) {
    if (!collapsed) setCollapsed(true)
      return true
    }
    return false;
}

return (
  <ConfigProvider
    theme={{
      algorithm: ThemePicker(),
    }}
  >
    <Layout>
      {authenticated && (
        <Sider trigger={null} collapsible collapsed={collapsed}>
          <Navigation authenticated={authenticated} isCollapsed={collapsed} setCollapsed={setCollapsed} />
        </Sider>
      )}
      <Layout>
        {authenticated && !isDocPage() && (
          <div style={{ padding: 0, height: compactUI ? '40px' : '48px', width: '100%'}}>
            <HeaderMenu />
          </div>
        )}
        <Content style={{ padding: isDocPage() ? 0 : 8, minHeight: 1200 }}>
          <Routes>
            <Route index element={<div>Home</div>} />

            <Route path="/test" element={<Testing />} />

            <Route path="/my-details" element={<MyPage />} />
            <Route path="/login" element={<Login />} />

            <Route path="/user/:id" element={<User />} />
            <Route path="/users" element={<UsersDetails />} />

            <Route path="/project/:id" element={<Project />} />
            <Route path="/projects" element={<Projects />} />
            <Route path="/create-project" element={<ProjectCreate />} />

            <Route path="/sub-projects" element={<SubProjects />} />
            <Route path="/create-sub-project" element={<SubProjectCreate />} />
            <Route path="/sub-project/:id" element={<SubProject />} />

            <Route path="/privileges" element={<Privileges />} />
            <Route path="/privilege/:id" element={<Privilege />} />

            <Route path="/expenses" element={<Expenses />} />
            <Route path="/create-expense" element={<ExpenseCreate />} />
            <Route path="/expense/:id" element={<Expense />} />

            <Route path="/incomes" element={<Incomes />} />
            {/* <Route path="/create-income" element={<IncomeCreate />} /> */}
            <Route path="/income/:id" element={<Income />} />

            <Route path="/external-companies" element={<ExternalCompanies />} />
            <Route path="/create-external-company" element={<ExternalCompanyCreate />} />
            <Route path="/external-company/:id" element={<ExternalCompany />} />

            <Route path="/products" element={<Products />} />

            <Route path="/invoice/:id" element={<Invoice />} />
            <Route path="/invoices" element={<Invoices />} />
            <Route path="/create-invoice" element={<CreateInvoice />} />

            <Route path="/invoice-item/:id" element={<InvoiceItem />} />
            <Route path="/invoice-items" element={<InvoiceItems />} />

            <Route path="/services" element={<ServiceOverviewPage />} />

            <Route path="/documentation" element={<Documentation />} />

          </Routes>
        </Content>
      </Layout>
    </Layout>
  </ConfigProvider>
  );
};

export default App;
