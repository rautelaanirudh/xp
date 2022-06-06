import React, { useContext, useEffect, useState } from "react";

import {
  EuiDragDropContext,
  EuiDraggable,
  EuiDroppable,
  EuiFlexGroup,
  EuiFlexItem,
  EuiForm,
  EuiFormRow,
  EuiGlobalToastList,
  EuiIcon,
  EuiPanel,
  EuiSpacer,
  euiDragDropMove,
  euiDragDropReorder,
} from "@elastic/eui";
import { FormLabelWithToolTip } from "@gojek/mlp-ui";
import isEqual from "lodash/isEqual";

import { Panel } from "components/panel/Panel";
import SegmenterContext from "providers/segmenters/context";
import { makeId } from "utils/helpers";

import { SegmenterCard } from "./SegmenterCard";

let toastId = 0;

const makeSegmenterItem = (item, selectedVariables) => {
  const newItem = {
    id: makeId(),
    name: item.name,
    variables: item.variables || [],
    selectedVariables: selectedVariables || [],
    isRequired: item.required || false,
  };
  // Auto-assign selected variables where the variables list has only 1 option
  if (!newItem.selectedVariables.length && newItem.variables.length === 1) {
    newItem.selectedVariables = newItem.variables[0];
  }
  return newItem;
};

const makeSegmenterCard = (
  { id, name, variables, selectedVariables, isRequired },
  idx,
  isExpandable,
  onChangeSelectedVariables,
  errors
) => {
  return (
    <EuiDraggable
      spacing="m"
      key={id}
      index={idx}
      draggableId={id}
      customDragHandle={true}
      hasInteractiveChildren={true}>
      {(provided, state) => (
        <SegmenterCard
          id={id}
          name={name}
          variables={variables}
          selectedVariables={selectedVariables}
          errors={(errors || {})[name] || ""}
          isRequired={isRequired}
          isDragging={state.isDragging}
          isExpandable={isExpandable}
          onChangeSelectedVariables={onChangeSelectedVariables(name)}
          dragHandleProps={provided.dragHandleProps}
        />
      )}
    </EuiDraggable>
  );
};

export const SegmenterPanel = ({ segmenters, onChange, errors = {} }) => {
  const [toasts, setToasts] = useState([]);
  const [availableSegmenters, setAvailableSegmenters] = useState([]);
  const [selectedSegmenters, setSelectedSegmenters] = useState(() => {
    // Init the selected segmenter names correctly, for the edit scenario. If not,
    // the useEffect hook will clear the data after the segmenterConfig is first loaded.
    return segmenters.names.map((name) => makeSegmenterItem({ name }));
  });
  const { segmenterConfig, isLoaded } = useContext(SegmenterContext);

  // This hook handles changes to the selected segmenter names.
  useEffect(() => {
    const selectedSegmenterNames = selectedSegmenters.map((s) => s.name);
    // Check if the number of segmenters or their order changed and update the properties individually.
    // Updating segmenters as a whole will cause the list to be re-initialized.
    // We want to avoid re-loading the entire list to retain the card state (such as if collapsed).
    if (!isEqual(selectedSegmenterNames, segmenters.names)) {
      onChange("segmenters.names")(selectedSegmenterNames);
      onChange("segmenters.variables")(
        selectedSegmenters.reduce((acc, e) => {
          acc[e.name] = e.selectedVariables;
          return acc;
        }, {})
      );
    }
  }, [selectedSegmenters, segmenters.names, onChange]);

  const onChangeSelectedVariables = (name) => (value) => {
    // Update the selected segmenters' local state
    const newSegmenters = selectedSegmenters.map((e) => {
      return e.name === name ? makeSegmenterItem(e, value) : e;
    });
    setSelectedSegmenters(newSegmenters);
    // Update the form state
    onChange(`segmenters.variables.${name}`)(value);
  };

  useEffect(() => {
    if (isLoaded) {
      let availableDraggableSegmenters = [];
      let selectedDraggableSegmenters = [];
      segmenterConfig.forEach((s) => {
        const item = makeSegmenterItem(s, segmenters.variables[s.name]);
        const isSelected = segmenters.names.includes(s.name);
        if (!isSelected && !item.required) {
          // Not used previously
          availableDraggableSegmenters.push(item);
        } else {
          // Retain the order
          const orderIdx = segmenters.names.indexOf(s.name);
          selectedDraggableSegmenters.splice(orderIdx, 0, item);
        }
      });
      setAvailableSegmenters(availableDraggableSegmenters);
      setSelectedSegmenters(selectedDraggableSegmenters);
    }
  }, [segmenterConfig, isLoaded, segmenters]);

  const onDragEnd = ({ source, destination }) => {
    const lists = {
      AVAILABLE_SEGMENTERS_DROPPABLE_AREA: availableSegmenters,
      SELECTED_SEGMENTERS_DROPPABLE_AREA: selectedSegmenters,
    };
    const actions = {
      AVAILABLE_SEGMENTERS_DROPPABLE_AREA: setAvailableSegmenters,
      SELECTED_SEGMENTERS_DROPPABLE_AREA: setSelectedSegmenters,
    };
    if (source && destination) {
      if (source.droppableId === destination.droppableId) {
        const items = euiDragDropReorder(
          lists[destination.droppableId],
          source.index,
          destination.index
        );
        actions[destination.droppableId](items);
      } else {
        const sourceId = source.droppableId;
        const destinationId = destination.droppableId;
        if (
          sourceId === "SELECTED_SEGMENTERS_DROPPABLE_AREA" &&
          lists[sourceId][source.index].isRequired
        ) {
          setToasts(
            toasts.concat({
              id: `toast${toastId++}`,
              title: "Oops, there was an error",
              color: "danger",
              iconType: "help",
              text: (
                <p>
                  <b>{lists[sourceId][source.index].name}</b> is a required
                  segmenter, it cannot be removed.
                </p>
              ),
            })
          );
          return;
        }
        const result = euiDragDropMove(
          lists[sourceId],
          lists[destinationId],
          source,
          destination
        );
        actions[sourceId](result[sourceId]);
        actions[destinationId](result[destinationId]);
      }
    }
  };

  const removeToast = (removedToast) => {
    setToasts(toasts.filter((toast) => toast.id !== removedToast.id));
  };

  const noSegmenterCards = (
    <EuiFlexGroup
      alignItems="center"
      justifyContent="spaceAround"
      gutterSize="none"
      style={{ height: "100%" }}>
      <EuiFlexItem grow={true}>
        <EuiIcon type="faceSad" />
      </EuiFlexItem>
    </EuiFlexGroup>
  );

  return (
    <>
      <EuiDragDropContext onDragEnd={onDragEnd}>
        <EuiPanel>
          <EuiForm>
            <EuiFormRow
              fullWidth
              label={
                <FormLabelWithToolTip
                  label="Drag and drop the Segmenters to be used in the project. The order defines the priority of the Segmenters."
                  content="Asterisk (*) Segmenters are required, and cannot be unselected."
                />
              }
              isInvalid={!!errors.segmenters?.names}
              error={errors.segmenters?.names}>
              <>
                <EuiSpacer />
                <EuiFlexGroup>
                  <EuiFlexItem>
                    <Panel title="Available Segmenters">
                      <EuiDroppable
                        droppableId="AVAILABLE_SEGMENTERS_DROPPABLE_AREA"
                        spacing="m"
                        grow={true}
                        ignoreContainerClipping={true}>
                        {availableSegmenters.length > 0
                          ? availableSegmenters.map((item, idx) =>
                              makeSegmenterCard(
                                item,
                                idx,
                                false,
                                onChangeSelectedVariables,
                                errors.segmenters?.variables
                              )
                            )
                          : noSegmenterCards}
                      </EuiDroppable>
                    </Panel>
                  </EuiFlexItem>
                  <EuiFlexItem>
                    <Panel title="Selected Segmenters">
                      <EuiDroppable
                        droppableId="SELECTED_SEGMENTERS_DROPPABLE_AREA"
                        spacing="m"
                        grow={true}
                        ignoreContainerClipping={true}>
                        {selectedSegmenters.length > 0
                          ? selectedSegmenters.map((item, idx) =>
                              makeSegmenterCard(
                                item,
                                idx,
                                true,
                                onChangeSelectedVariables,
                                errors.segmenters?.variables
                              )
                            )
                          : noSegmenterCards}
                      </EuiDroppable>
                    </Panel>
                  </EuiFlexItem>
                </EuiFlexGroup>
              </>
            </EuiFormRow>
            <EuiSpacer size="m" />
          </EuiForm>
        </EuiPanel>
      </EuiDragDropContext>
      <EuiGlobalToastList
        toasts={toasts}
        dismissToast={removeToast}
        toastLifeTimeMs={3000}
      />
    </>
  );
};