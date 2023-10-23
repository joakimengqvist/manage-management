import { useSelector } from 'react-redux';
import Income from '../components/economics/incomes/income';
import { PRIVILEGES } from '../enums/privileges';
import { hasPrivilege } from '../helpers/hasPrivileges';
import { State } from '../types/state';

const PrivilegeDetails: React.FC = () => {
    const userPrivileges = useSelector((state : State) => state.user.privileges)

    if (!hasPrivilege(userPrivileges, PRIVILEGES.economics_read)) return null;
    
    return (
        <div style={{padding: '12px 8px'}}>
            <Income />
        </div>
    )

}

export default PrivilegeDetails;