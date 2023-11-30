import CreateProjectInvoice from '../../components/invoice/invoice/createInvoice';
import { PRIVILEGES } from '../../enums/privileges';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { useGetLoggedInUserPrivileges } from '../../hooks/useGetLoggedInUserPrivileges';

const CreateInvoice = () => {
    const userPrivileges = useGetLoggedInUserPrivileges();

    if (!hasPrivilege(userPrivileges, PRIVILEGES.economics_write)) return null;
    
    return <CreateProjectInvoice />
}

export default CreateInvoice;