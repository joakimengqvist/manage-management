import { Row, Col } from 'antd';
import { useSelector } from 'react-redux';
import { State } from '../../types/state';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { PRIVILEGES } from '../../enums/privileges';
import CreateInvoiceItem from '../../components/invoice/invoiceItem/CreateInvoiceItem';
import InvoiceIteams from '../../components/invoice/invoiceItem/InvoiceItems';

const InvoiceItems: React.FC = () => {
    const userPrivileges = useSelector((state : State) => state.user.privileges)
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