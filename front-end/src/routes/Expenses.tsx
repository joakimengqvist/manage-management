/* eslint-disable @typescript-eslint/no-explicit-any */
import { Button, Select } from 'antd';
import { useState } from 'react';
import { useSelector } from 'react-redux';
import { useNavigate } from 'react-router-dom';
import Expenses from '../components/economics/expenses/expenses';
import { State } from '../types/state';

const ExpensesRoute: React.FC = () => {
    const navigate = useNavigate();
    const projects = useSelector((state: State) => state.application.projects);
    const [project, setProject] = useState('all');

    if (!projects) {
        return null
    }

    const projectOptions = [{label: 'All projects', value: 'all'}]
    projects.forEach(project => {
        projectOptions.push({
            value: project.id,
            label: project.name,
        })
    });

    const onSelectProject = (value: any) => setProject(value);
    
    return (
        <div style={{padding: '12px 8px'}}>
            <div style={{padding: '4px'}}>
                <div style={{display: 'flex', justifyContent: 'space-between', paddingBottom: '12px'}}>
                <Select
                        defaultValue={projectOptions[0].value}
                        style={{ width: 300 }}
                        options={projectOptions}
                        onSelect={onSelectProject}
                />
                <Button type="primary" onClick={() => navigate("/create-expense")}>Create new expense</Button>
                </div>
                <Expenses project={project} />
            </div>
        </div>
    )

}

export default ExpensesRoute;