/* eslint-disable @typescript-eslint/no-explicit-any */
import { Button } from 'antd';
import { useNavigate } from 'react-router-dom';
import Invoices from '../../components/invoice/invoice/invoices';
import { PRIVILEGES } from '../../enums/privileges';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { useGetLoggedInUserPrivileges } from '../../hooks/useGetLoggedInUserPrivileges';

const InvoiceDetails = () => {
    const navigate = useNavigate();
    const userPrivileges = useGetLoggedInUserPrivileges();

    if (!hasPrivilege(userPrivileges, PRIVILEGES.invoice_read)) return null;

    return (<>
        <div style={{display: 'flex', justifyContent: 'flex-end', paddingBottom: '8px'}}>
            <Button onClick={() => navigate("/create-invoice")}>Create new invoice</Button>
        </div>
        <Invoices />
    </>)
}

export default InvoiceDetails;