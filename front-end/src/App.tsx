/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable react-hooks/exhaustive-deps */
import { useState, useEffect } from 'react';
import { Route, Routes, useNavigate } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import './App.css';
import HeaderMenu from './components/header/HeaderMenu';
import Navigation from './components/navigation/Navigation';
import Login from './routes/login';
import UsersDetails from './routes/user/Users';
import User from './routes/user/User';
import Projects from './routes/project/Projects';
import Project from './routes/project/Project';
import Privileges from './routes/privilege/Privileges';
import Privilege from './routes/privilege/Privilege';
import { ConfigProvider, Layout, theme } from 'antd';
import { getAllPrivileges } from './api/privileges/getAll';
import { clearData, fetchExternalCompanies, fetchInvoiceItems, fetchInvoices, fetchPrivileges, fetchProducts, fetchProjects, fetchSubProjects, fetchUsers, initiateApplicationData } from './redux/applicationDataSlice';
import { fetchUserSettings, initiateUser } from './redux/userDataSlice'
import { getAllUsers } from './api/users/getAll';
import { getAllProjects } from './api/projects/getAll';
import { State } from './types/state';
import { hasPrivilege } from './helpers/hasPrivileges';
import { PRIVILEGES } from './enums/privileges';
import MyPage from './routes/myPage';
import ExpenseCreate from './routes/expense/ExpenseCreate';
import Expense from './routes/expense/Expense';
import Expenses from './routes/expense/Expenses';
import ServiceOverview from './routes/ServiceOverview';
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

const { Sider, Content } = Layout;

const App: React.FC = () => {
  const dispatch = useDispatch();
  const navigate = useNavigate();
  const [isInitiated, setIsInitiated] = useState(false);
  const authenticated = useSelector((state : State) => state.user.authenticated);
  const loggedInUserId = useSelector((state : State) => state.user.id);
  const UserSettings = useSelector((state : State) => state.user.settings)
  const userPrivileges = useSelector((state : State) => state.user.privileges)
  const [collapsed, setCollapsed] = useState(false);

  if (!authenticated) {
    dispatch(clearData())
    navigate('/login')
  }

  useEffect(() => {
    if (!isInitiated) {
      dispatch(initiateUser());
      dispatch(initiateApplicationData());
      setIsInitiated(true);
    }

    if (loggedInUserId && authenticated) {
      getUserSettingsByUserId(loggedInUserId, loggedInUserId).then(response => dispatch(fetchUserSettings(response.data))).catch(() => {})

      if (hasPrivilege(userPrivileges, PRIVILEGES.privilege_read)) {
        getAllPrivileges(loggedInUserId).then(response => dispatch(fetchPrivileges(response.data))).catch(() => {})
      }
      if (hasPrivilege(userPrivileges, PRIVILEGES.user_read)) {
        getAllUsers(loggedInUserId).then(response => dispatch(fetchUsers(response.data))).catch(() => {})
      }
      if (hasPrivilege(userPrivileges, PRIVILEGES.project_read)) {
        getAllProjects(loggedInUserId).then(response => dispatch(fetchProjects(response.data))).catch(() => {})
      }
      if (hasPrivilege(userPrivileges, PRIVILEGES.sub_project_read)) {
        getAllSubProjects(loggedInUserId).then(response => dispatch(fetchSubProjects(response.data))).catch(() => {})
      }
      if (hasPrivilege(userPrivileges, PRIVILEGES.external_company_read)) {
        getAllExternalCompanies(loggedInUserId).then(response => dispatch(fetchExternalCompanies(response.data))).catch(() => {})
      }
      if (hasPrivilege(userPrivileges, PRIVILEGES.product_read)) {
        getAllProducts(loggedInUserId).then(response => dispatch(fetchProducts(response.data))).catch(() => {})
      }
      if (hasPrivilege(userPrivileges, PRIVILEGES.invoice_read)) {
        getAllInvoices(loggedInUserId).then(response => dispatch(fetchInvoices(response.data))).catch(() => {})
      }
      if (hasPrivilege(userPrivileges, PRIVILEGES.invoice_read)) {
        getAllInvoiceItems(loggedInUserId).then(response => dispatch(fetchInvoiceItems(response.data))).catch(() => {})
      }
    }
}, [loggedInUserId])

const ThemePicker = () => {
  if (UserSettings.compact_ui) {
    if (UserSettings.dark_theme) {
      return [theme.darkAlgorithm, theme.compactAlgorithm]
    } else {
      return [theme.compactAlgorithm]
    }
  }

  return [UserSettings.dark_theme ? theme.darkAlgorithm : theme.defaultAlgorithm]
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
          {authenticated && (
            <div style={{ padding: 0, height: '48px', width: '100%'}}>
              <HeaderMenu />
            </div>
          )}
          <Content style={{ padding: 8, minHeight: 1200 }}>
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



              <Route path="/services" element={<ServiceOverview />} />

            </Routes>
          </Content>
        </Layout>
      </Layout>
    </ConfigProvider>
  );
};

export default App;
