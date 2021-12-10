import React, { useState, useEffect, ReactElement } from 'react';
import { Button, ButtonVariant, Flex, FlexItem, SelectOption } from '@patternfly/react-core';

import SelectSingle from 'Components/SelectSingle';
import FormLabelGroup from 'Components/PatternFly/FormLabelGroup';
import { fetchAccessScopes, AccessScope } from 'services/AccessScopesService';

type ResourceScopeSelectionProps = {
    scopeId: string;
    setFieldValue: (field: string, value: any, shouldValidate?: boolean | undefined) => void;
};

function ResourceScopeSelection({
    scopeId,
    setFieldValue,
}: ResourceScopeSelectionProps): ReactElement {
    const [resourceScopes, setResourceScopes] = useState<AccessScope[]>([]);

    function getScopes(): void {
        fetchAccessScopes()
            .then((response) => {
                const resourceScopesList = response || [];
                const filteredScopes = resourceScopesList.filter(
                    (scope) => scope.id !== 'io.stackrox.authz.accessscope.denyall'
                );
                setResourceScopes(filteredScopes);
            })
            .catch(() => {
                // TODO display message when there is a place for minor errors
            });
    }

    useEffect(() => {
        getScopes();
    }, []);

    function onScopeChange(_id, selection) {
        setFieldValue('scopeId', selection);
    }

    return (
        <Flex alignItems={{ default: 'alignItemsFlexEnd' }}>
            <FlexItem>
                <FormLabelGroup
                    className="pf-u-mb-md"
                    isRequired
                    label="Configure resource scope"
                    fieldId="scopeId"
                    touched={{}}
                    errors={{}}
                >
                    <SelectSingle
                        id="scopeId"
                        value={scopeId}
                        handleSelect={onScopeChange}
                        isDisabled={false}
                        placeholderText="Select a scope"
                    >
                        {resourceScopes.map(({ id, name, description }) => (
                            <SelectOption key={id} value={id} description={description}>
                                {name}
                            </SelectOption>
                        ))}
                    </SelectSingle>
                </FormLabelGroup>
            </FlexItem>
            <FlexItem>
                <Button className="pf-u-mb-md" variant={ButtonVariant.secondary}>
                    Create resource scope
                </Button>
            </FlexItem>
        </Flex>
    );
}

export default ResourceScopeSelection;