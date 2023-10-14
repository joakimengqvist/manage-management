import { useSelector } from 'react-redux';
import CreateProjectIncome from '../components/economics/incomes/createProjectIncome';
import { hasPrivilege } from '../helpers/hasPrivileges';
import { State } from '../types/state';

const CreateIncome: React.FC = () => {
    const userPrivileges = useSelector((state : State) => state.user.privileges);

    if (!hasPrivilege(userPrivileges, 'economics_write')) return null;
    
    return (
        <div style={{padding: '12px 8px'}}>
            <div style={{padding: '4px'}}>
                <CreateProjectIncome />
            </div>
        </div>
    )

}

export default CreateIncome;