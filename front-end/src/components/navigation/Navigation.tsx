import { useMemo } from 'react';
import { useNavigate } from "react-router-dom";
import {
  BarcodeOutlined,
  MenuFoldOutlined,
  MenuUnfoldOutlined,
  IdcardOutlined,
  TeamOutlined,
  UnlockOutlined,
  ProjectOutlined,
  FundProjectionScreenOutlined,
  ExceptionOutlined,
  DollarOutlined,
  TranslationOutlined,
  FundOutlined,
 // ClusterOutlined,
  BuildOutlined,
  BankOutlined,
  ReadOutlined,
} from '@ant-design/icons';
import type { MenuProps } from 'antd';
import { Menu } from 'antd';

type MenuItem = Required<MenuProps>['items'][number];

function getItem(
  label: React.ReactNode,
  key: React.Key,
  icon?: React.ReactNode,
  onClick?: () => void,
  children?: MenuItem[],
  type?: 'group',
): MenuItem {
  return {
    onClick,
    key,
    icon,
    children,
    label,
    type,
  } as MenuItem;
}

const Navigation = (props : {
  isCollapsed: boolean,
  setCollapsed: (isCollapsed : boolean) => void,
  authenticated: boolean,
}) => {
    const { isCollapsed, setCollapsed } = props;
    const navigate = useNavigate();

  const itemsLoggedIn: MenuItem[] = useMemo(() => ([
    getItem(isCollapsed ? 'Expand' : 'Minimize', '0', isCollapsed ? <MenuUnfoldOutlined /> : <MenuFoldOutlined />, () => setCollapsed(!isCollapsed)),
    getItem('People', '1', <TeamOutlined />, () => {}, [
      getItem('Users', 'sub1-1', <IdcardOutlined />, () => navigate("/users")),
      getItem('Privileges', 'sub1-2', <UnlockOutlined />, () => navigate("/privileges")),
    ]),
    getItem('Projects', '2', <ProjectOutlined />, () => {}, [
      getItem('Projects', 'sub2-1', <FundProjectionScreenOutlined />, () => navigate("/projects")),
      getItem('Sub Projects', 'sub2-2', <BuildOutlined />, () => navigate("/sub-projects")),
    ]),
    getItem('Products', '3', <BarcodeOutlined />, () => navigate("/products")),
    getItem('Economics', '4', <DollarOutlined />, () => {}, [
      getItem('Invoice items', 'sub4-1', <TranslationOutlined />, () => navigate("/invoice-items")),
      getItem('Invoices', 'sub4-2', <ExceptionOutlined />, () => navigate("/invoices")),
      getItem('Incomes', 'sub4-3', <FundOutlined />, () => navigate("/incomes")),
      /* getItem('Expenses', 'sub3-2', <ReconciliationOutlined />, () => navigate("/expenses")), */
  ]),
    getItem('External companies', '5', <BankOutlined />, () => navigate("/external-companies")),
    getItem('Frontend docs', '6', <ReadOutlined />, () => navigate("/documentation")),
    // getItem('Services', '6', <ClusterOutlined />, () => navigate("/services")),
// eslint-disable-next-line react-hooks/exhaustive-deps
]), [isCollapsed])



  return (
      <Menu
        defaultSelectedKeys={['1']}
        defaultOpenKeys={['sub1']}
        mode="inline"
        inlineCollapsed={isCollapsed}
        items={itemsLoggedIn}
        style={{height: '100%'}}
        inlineIndent={10}
      />
  );
};

export default Navigation;