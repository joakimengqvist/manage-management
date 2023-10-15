import { useSelector } from 'react-redux';
import CreateExternalCompany from '../components/externalCompanies/CreateExternalCompany';
import { PRIVILEGES } from '../enums/privileges';
import { hasPrivilege } from '../helpers/hasPrivileges';
import { State } from '../types/state';

const CreateExpense: React.FC = () => {
    const userPrivileges = useSelector((state : State) => state.user.privileges);

    if (!hasPrivilege(userPrivileges, PRIVILEGES.external_company_write)) return null;
    
    return (
        <div style={{padding: '12px 8px'}}>
            <div style={{padding: '4px'}}>
                <CreateExternalCompany />
            </div>
        </div>
    )

}

export default CreateExpense;