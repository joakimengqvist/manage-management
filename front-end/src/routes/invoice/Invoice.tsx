import { useSelector } from 'react-redux';
import Invoice from '../../components/invoice/invoice/invoice';
import { PRIVILEGES } from '../../enums/privileges';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { State } from '../../interfaces/state';

const InvoiceDetails = () => {
    const userPrivileges = useSelector((state : State) => state.user.privileges)

    if (!hasPrivilege(userPrivileges, PRIVILEGES.invoice_read)) return null;
    
    return <Invoice />
}

export default InvoiceDetails;