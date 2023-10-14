import { useSelector } from 'react-redux';
import Income from '../components/economics/incomes/income';
import { hasPrivilege } from '../helpers/hasPrivileges';
import { State } from '../types/state';

const PrivilegeDetails: React.FC = () => {
    const userPrivileges = useSelector((state : State) => state.user.privileges)

    if (!hasPrivilege(userPrivileges, 'economics_read')) return null;
    
    return (
        <div style={{padding: '12px 8px'}}>
            <div style={{padding: '4px'}}>
                <Income />
            </div>
        </div>
    )

}

export default PrivilegeDetails;