import { useSelector } from 'react-redux';
import CreateProjectExpense from '../components/economics/expenses/createProjectExpense';
import { hasPrivilege } from '../helpers/hasPrivileges';
import { State } from '../types/state';

const CreateExpense: React.FC = () => {
    const userPrivileges = useSelector((state : State) => state.user.privileges);

    if (!hasPrivilege(userPrivileges, 'economics_write')) return null;
    
    return (
        <div style={{padding: '12px 8px'}}>
            <div style={{padding: '4px'}}>
                <CreateProjectExpense />
            </div>
        </div>
    )

}

export default CreateExpense;