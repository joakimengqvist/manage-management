import { useSelector } from 'react-redux';
import CreateProjectInvoice from '../../components/invoice/invoice/createInvoice';
import { PRIVILEGES } from '../../enums/privileges';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { State } from '../../interfaces/state';

const CreateInvoice = () => {
    const userPrivileges = useSelector((state : State) => state.user.privileges);

    if (!hasPrivilege(userPrivileges, PRIVILEGES.economics_write)) return null;
    
    return <CreateProjectInvoice />
}

export default CreateInvoice;