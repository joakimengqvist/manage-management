import { useSelector } from 'react-redux';
import CreateExternalCompany from '../../components/externalCompanies/createExternalCompany';
import { PRIVILEGES } from '../../enums/privileges';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { State } from '../../interfaces/state';

const CreateExpense = () => {
    const userPrivileges = useSelector((state : State) => state.user.privileges);

    if (!hasPrivilege(userPrivileges, PRIVILEGES.external_company_write)) return null;
    
    return <CreateExternalCompany />
}

export default CreateExpense;