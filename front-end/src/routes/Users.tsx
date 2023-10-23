import Users from '../components/users/Users';
import CreateUser from '../components/users/CreateUser';
import { Row, Col } from 'antd';
import { useSelector } from 'react-redux';
import { State } from '../types/state';
import { hasPrivilege } from '../helpers/hasPrivileges';
import { PRIVILEGES } from '../enums/privileges';

const UsersDetails: React.FC = () => {
    const userPrivileges = useSelector((state : State) => state.user.privileges)
    return (
        <div style={{padding: '12px 8px'}}>
            <Row>
                <Col span={16}>
                    <div style={{paddingRight: '12px'}}>
                        {hasPrivilege(userPrivileges, PRIVILEGES.user_read) && <Users />}
                    </div>
                </Col>
                <Col span={8}>
                    {hasPrivilege(userPrivileges, PRIVILEGES.user_write) && <CreateUser />}
                </Col>
            </Row>
        </div>
    )

}

export default UsersDetails;