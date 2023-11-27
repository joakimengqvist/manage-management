import { useSelector } from 'react-redux';
import InvoiceItem from '../../components/invoice/invoiceItem/InvoiceItem';
import { PRIVILEGES } from '../../enums/privileges';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { State } from '../../interfaces/state';

const InvoiceItemDetails = () => {
    const userPrivileges = useSelector((state : State) => state.user.privileges)

    if (!hasPrivilege(userPrivileges, PRIVILEGES.invoice_read)) return null;
    
    return <InvoiceItem />
}

export default InvoiceItemDetails;