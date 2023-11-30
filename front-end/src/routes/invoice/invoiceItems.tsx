import { Row, Col } from 'antd';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { PRIVILEGES } from '../../enums/privileges';
import CreateInvoiceItem from '../../components/invoice/invoiceItem/CreateInvoiceItem';
import InvoiceIteams from '../../components/invoice/invoiceItem/InvoiceItems';
import { useGetLoggedInUserPrivileges } from '../../hooks/useGetLoggedInUserPrivileges';

const InvoiceItems = () => {
    const userPrivileges = useGetLoggedInUserPrivileges();
    return (
        <Row>
            <Col span={16}>
                <div style={{paddingRight: '8px'}}>
                    {hasPrivilege(userPrivileges, PRIVILEGES.invoice_read) && <InvoiceIteams />}
                </div>
            </Col>
            <Col span={8}>
                {hasPrivilege(userPrivileges, PRIVILEGES.invoice_write) && <CreateInvoiceItem />}
            </Col>
        </Row>
    )

}

export default InvoiceItems;