import Privilege from '../../components/privileges/Privilege';
import { Row, Col } from 'antd';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { PRIVILEGES } from '../../enums/privileges';
import { useGetLoggedInUserPrivileges } from '../../hooks/useGetLoggedInUserPrivileges';

const PrivilegeDetails = () => {
    const userPrivileges = useGetLoggedInUserPrivileges();
    return (
        <Row>
            <Col span={8}>
            {hasPrivilege(userPrivileges, PRIVILEGES.privilege_read) &&
                <Privilege />
            }
            </Col>
            <Col span={16}>
            </Col>
        </Row>
    )

}

export default PrivilegeDetails;